FROM golang:1.23.3-alpine3.20

RUN \
    apk --no-cache add \
        git \
        bash \
        coreutils \
        ca-certificates \
    && mkdir -p /go/src/github.com/PyYoshi/pubsub-gateway/cmd/gateway \
    && go get -u -v \
        github.com/oxequa/realize

# ENV GO111MODULE=on

WORKDIR /go/src/github.com/PyYoshi/pubsub-gateway

CMD ["bash", "./cmd/gateway/dev.sh"]
