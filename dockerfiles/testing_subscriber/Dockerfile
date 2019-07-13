FROM golang:1.12-alpine3.9


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

CMD ["bash", "./cmd/testing_subscriber/dev.sh"]
