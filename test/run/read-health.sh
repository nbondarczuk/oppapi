#!/bin/bash

HOST=localhost
PORT=8080
URL=http://${HOST}:${PORT}/health
CMD="curl -s -v --trace-ascii out.txt $URL"
echo -n Running command: $CMD " - result: "
eval $CMD | jq
