FROM golang:1.12-alpine3.9 AS builder

WORKDIR /go/src/github.com/PyYoshi/pubsub-gateway

COPY . .

RUN \
  mkdir -p /go/src/github.com/PyYoshi/pubsub-gateway/dist \
  && cd cmd/gateway/ \
  && go build -o /go/src/github.com/PyYoshi/pubsub-gateway/dist/gateway-server

FROM alpine:3.9

RUN \
    apk --no-cache add \
        bash \
        coreutils \
        ca-certificates \
    && mkdir -p /opt/pubsub-gateway/bin

COPY --from=builder /go/src/github.com/PyYoshi/pubsub-gateway/dist/gateway-server /opt/pubsub-gateway/bin/gateway-server
COPY --from=builder /go/src/github.com/PyYoshi/pubsub-gateway/gen/http/openapi.yaml /opt/pubsub-gateway/openapi.yaml
COPY --from=builder /go/src/github.com/PyYoshi/pubsub-gateway/gen/http/openapi.json /opt/pubsub-gateway/openapi.json

RUN \
  chmod +x /opt/pubsub-gateway/bin/gateway-server

CMD ["/opt/pubsub-gateway/bin/gateway-server"]
