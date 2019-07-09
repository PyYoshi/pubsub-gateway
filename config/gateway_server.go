package config

import (
	"github.com/kelseyhightower/envconfig"
)

// GatewayServer GatewayServer用設定情報
type GatewayServer struct {
	Debug                bool         `envconfig:"DEBUG" default:"false"`
	BindAddress          string       `envconfig:"BIND_ADDRESS" default:"0.0.0.0:8089" required:"true"`
	GoogleServiceAccount Base64String `envconfig:"GOOGLE_SERVICE_ACCOUNT_BASE64" required:"true"`
	GoogleProjectID      string       `envconfig:"GOOGLE_PROJECT_ID" required:"true"`
}

// NewGatewayServer create a GatewayServer instance
func NewGatewayServer() (*GatewayServer, error) {
	cfg := &GatewayServer{}
	err := envconfig.Process("GATEWAY_SERVER", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
