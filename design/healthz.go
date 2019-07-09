package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("healthz", func() {
	dsl.Description("ヘルスチェックサービス")

	dsl.HTTP(func() {
		dsl.Path("/healthz")
	})

	dsl.Method("readiness", func() {
		dsl.Description("サーバが準備完了状態でいるかチェックするエンドポイント")

		dsl.HTTP(func() {
			dsl.GET("/readiness")
			dsl.Response(dsl.StatusOK)
			dsl.Response(dsl.StatusInternalServerError)
		})
	})

	dsl.Method("liveness", func() {
		dsl.Description("サーバが生きているかチェックするエンドポイント")

		dsl.HTTP(func() {
			dsl.GET("/liveness")
			dsl.Response(dsl.StatusOK)
			dsl.Response(dsl.StatusInternalServerError)
		})
	})
})
