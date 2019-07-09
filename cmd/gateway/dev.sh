#!/bin/bash

bash /go/src/github.com/PyYoshi/cloud_pubsub_gateway/dockerfiles/wait-for-it.sh \
    -h pubsub_emulator \
    -p 8086 \
    -t 600 -- echo "wait-for-it.sh: 🍣💨🍣💨🍣💨"

cd cmd/gateway
realize start
