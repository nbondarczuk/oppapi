#!/bin/bash

HOST=localhost
PORT=8080
URL=http://${HOST}:${PORT}/payment/66c8f8f99376e5e86e9d0ea9
CMD="curl $URL"
echo -n Running command: $CMD " - result: "
eval $CMD
