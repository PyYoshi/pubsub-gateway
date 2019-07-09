package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"cloud.google.com/go/pubsub"
	pubusubgateway "github.com/PyYoshi/pubsub-gateway"
	"github.com/PyYoshi/pubsub-gateway/config"
	gcp "github.com/PyYoshi/pubsub-gateway/gen/gcp"
	healthz "github.com/PyYoshi/pubsub-gateway/gen/healthz"
	log "github.com/PyYoshi/pubsub-gateway/gen/log"
	"google.golang.org/api/option"
)

func main() {
	cfg, err := config.NewGatewayServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	var pubsubCli *pubsub.Client
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		pubsubCli, err = pubsub.NewClient(
			context.Background(),
			cfg.GoogleProjectID,
			option.WithCredentialsJSON([]byte(cfg.GoogleServiceAccount)),
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	} else {
		pubsubCli, err = pubsub.NewClient(
			context.Background(),
			cfg.GoogleProjectID,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	}

	defer func() {
		pbcErr := pubsubCli.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", pbcErr)
			os.Exit(1)
		}
	}()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New("pubusubgateway", !cfg.Debug)
	}

	// Initialize the services.
	var (
		gcpSvc     gcp.Service
		healthzSvc healthz.Service
	)
	{
		gcpSvc = pubusubgateway.NewGcp(logger, cfg, pubsubCli)
		healthzSvc = pubusubgateway.NewHealthz(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		gcpEndpoints     *gcp.Endpoints
		healthzEndpoints *healthz.Endpoints
	)
	{
		gcpEndpoints = gcp.NewEndpoints(gcpSvc)
		healthzEndpoints = healthz.NewEndpoints(healthzSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctxServer, cancelServer := context.WithCancel(context.Background())
	{
		addr := "http://localhost:8088"
		u, err := url.Parse(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
			os.Exit(1)
		}

		if cfg.BindAddress != "" {
			u.Host = cfg.BindAddress
		}

		handleHTTPServer(
			ctxServer,
			u,
			gcpEndpoints,
			healthzEndpoints,
			&wg,
			errc,
			logger,
			cfg,
		)
	}

	// Wait for signal.
	logger.Infof("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancelServer()

	wg.Wait()
	logger.Info("exited")
}
