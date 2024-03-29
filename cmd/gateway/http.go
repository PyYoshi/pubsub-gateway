package main

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/PyYoshi/pubsub-gateway/config"
	gcp "github.com/PyYoshi/pubsub-gateway/gen/gcp"
	healthz "github.com/PyYoshi/pubsub-gateway/gen/healthz"
	gcpsvr "github.com/PyYoshi/pubsub-gateway/gen/http/gcp/server"
	healthzsvr "github.com/PyYoshi/pubsub-gateway/gen/http/healthz/server"

	// swaggersvr "github.com/PyYoshi/pubsub-gateway/gen/http/swagger/server"
	log "github.com/PyYoshi/pubsub-gateway/gen/log"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(
	ctxServer context.Context,
	u *url.URL,
	gcpEndpoints *gcp.Endpoints,
	healthzEndpoints *healthz.Endpoints,
	wg *sync.WaitGroup,
	errc chan error,
	logger *log.Logger,
	cfg *config.GatewayServer,
) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		gcpServer     *gcpsvr.Server
		healthzServer *healthzsvr.Server
		// swaggerServer *swaggersvr.Server
	)
	{
		eh := errorHandler(logger)
		gcpServer = gcpsvr.New(gcpEndpoints, mux, dec, enc, eh, nil)
		healthzServer = healthzsvr.New(healthzEndpoints, mux, dec, enc, eh, nil)
		// swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh)
	}
	// Configure the mux.
	gcpsvr.Mount(mux, gcpServer)
	healthzsvr.Mount(mux, healthzServer)
	// swaggersvr.Mount(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if cfg.Debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range gcpServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range healthzServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	// for _, m := range swaggerServer.Mounts {
	// 	logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	// }

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Infof("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctxServer.Done()
		logger.Infof("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctxServer, cancel := context.WithTimeout(ctxServer, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctxServer)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.With(zap.String("id", id)).Error(err.Error())
	}
}
