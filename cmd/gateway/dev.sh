#!/bin/bash

bash /go/src/github.com/PyYoshi/pubsub-gateway/dockerfiles/wait-for-it.sh \
    -h pubsub_emulator \
    -p 8086 \
    -t 600 -- echo "wait-for-it.sh: 🍣💨🍣💨🍣💨"

cd cmd/gateway
realize start
