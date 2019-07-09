package design

import (
	"goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = dsl.API("pubusub-gateway", func() {
	dsl.Title("Pub/Sub Gateway")
	dsl.Description("HTTPを通してPub/Subが使えるようにするゲートウェイサーバ")

	dsl.Server("gateway", func() {
		dsl.Host("development", func() {
			dsl.Description("Development hosts")
			dsl.URI("http://localhost:8088")
		})
	})

	dsl.HTTP(func() {
		dsl.Path("/v1")
		dsl.Consumes("application/json")
		dsl.Produces("application/json")
	})
})
