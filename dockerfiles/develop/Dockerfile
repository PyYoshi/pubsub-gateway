FROM golang:1.12-alpine3.9

RUN \
    apk --no-cache add \
        git \
        bash \
    && mkdir -p /go/src/github.com/PyYoshi/cloud_pubsub_gateway/cmd/cloud_pubsub_gateway \
    && go get -u -v \
        github.com/oxequa/realize

# ENV GO111MODULE=on

WORKDIR /go/src/github.com/PyYoshi/cloud_pubsub_gateway

CMD ["bash", "./cmd/gateway/dev.sh"]
