#!/bin/bash

PORT=8080
HOST=localhost
URL=http://${HOST}:${PORT}/payment
HEADER1="\"Content-Type: application/json\""
HEADER2="\"X-API-KEY: R6bSXS4pfo7bnI0zIdMqiA=\""
CMD="curl -s -v --trace-ascii out.txt -H $HEADER1 -H $HEADER2 --data-binary "@examples/create-payment-request.json" $URL"
echo -n Request:
jq <examples/create-payment-request.json
echo Running command: $CMD
eval $CMD >out.json 2>out.err
echo -n Response:
jq <out.json
