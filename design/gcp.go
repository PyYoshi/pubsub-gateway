package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("gcp", func() {
	dsl.Description("Google Cloud Pub/Sub向けサービス")

	dsl.HTTP(func() {
		dsl.Path("/gcp")
	})

	dsl.Method("publish", func() {
		dsl.Description("httpプロトコルを経由してCloud Pub/Subにメッセージを渡すエンドポイント")

		dsl.Payload(GcpPublishRequestType)
		dsl.Result(GcpPublishResponseResultType)

		dsl.HTTP(func() {
			dsl.POST("/publish/{topic}")
			dsl.Response(dsl.StatusOK)
		})
	})
})

var GcpPublishRequestType = dsl.Type("GcpPublishRequestType", func() {
	dsl.Description("HTTPリクエストでCloud Pub/Subを実行するためのモデルのType")
	dsl.Attribute("topic", dsl.String, func() {
		dsl.Description("Cloud Pub/SubのTopic")
		dsl.MinLength(1)
	})
	dsl.Attribute("message", dsl.String, func() {
		dsl.Description("Cloud Pub/SubのメッセージBody")
		dsl.MinLength(1)
	})
	dsl.Required("topic", "message")
})

var GcpPublishResponseType = dsl.Type("GcpPublishResponseType", func() {
	dsl.Description("Cloud Pub/SubでレスポンスされるモデルのType")
	dsl.Attribute("id", dsl.String, "message id")
	dsl.Required("id")
})

var GcpPublishResponseResultType = dsl.ResultType("application/vnd.pubsub-gateway.gcp-publish", func() {
	dsl.Description("Cloud Pub/SubでレスポンスされるモデルのResultType")
	dsl.Reference(GcpPublishResponseType)

	dsl.Attributes(func() {
		dsl.Field(1, "id")
	})

	dsl.View("default", func() {
		dsl.Attribute("id")
	})

	dsl.Required("id")
})
