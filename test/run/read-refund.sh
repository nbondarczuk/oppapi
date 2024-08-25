#!/bin/bash

HOST=localhost
PORT=8080
HEADER1="X-API-KEY: R6bSXS4pfo7bnI0zIdMqiA="
URL=http://${HOST}:${PORT}/refund/$1
CMD="curl -s -v --trace-ascii out.txt -H $HEADER1 $URL"
echo -n Running command: $CMD " - result: "
eval $CMD | jq
