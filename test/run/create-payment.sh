#!/bin/bash

PORT=8080
HOST=localhost
URL=http://${HOST}:${PORT}/payment
HEADER="\"Content-Type: application/json\""
CMD="curl -s -v --trace-ascii out.txt -H $HEADER --data-binary "@examples/create-payment-request.json" $URL"
echo -n Request:
jq <examples/create-payment-request.json
echo Running command: $CMD
eval $CMD >out.json 2>out.err
echo -n Response:
jq <out.json
