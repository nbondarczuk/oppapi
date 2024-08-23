FROM golang:1.22.3-alpine as builder

ARG VERSION

ENV GOPATH /go

ENV GOCACHE /go/caches/go-build

RUN --mount=type=cache,target="/go/caches/go-build" apk add git alpine-sdk

ADD . /workpath

WORKDIR /workpath

# build the source
RUN --mount=type=cache,target="/go/caches/go-build" make go/tidy
RUN --mount=type=cache,target="/go/caches/go-build" make

# use a minimal alpine image
FROM alpine:latest

# add ca-certificates in case you need them
RUN --mount=type=cache,target="/go/caches/go-build" apk update && apk add ca-certificates curl bash

# set working directory
WORKDIR /work

COPY --from=builder /workpath/bin/oppapi /work/oppapi
COPY --from=builder /workpath/config/config.yaml /work/config/config.yaml

USER 1001
EXPOSE 8080

ARG GODEBUG

# run the binary
ENTRYPOINT ["./oppapi"]
