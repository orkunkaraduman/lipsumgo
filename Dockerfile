# syntax=docker.io/docker/dockerfile:1.4.3

FROM golang:1.18-alpine3.16 AS base

RUN apk upgrade --no-cache && apk add --no-cache \
    make git

FROM base AS builder

WORKDIR /src
COPY . .
RUN \
    --mount=type=cache,target=/root/go \
    --mount=type=cache,target=/root/.cache/go-build \
    make build
RUN mv target/ /app/
RUN cp -a docker/entrypoint.sh /app/

FROM alpine:3.16

RUN apk upgrade --no-cache && apk add --no-cache \
    bash openssl ca-certificates

COPY --from=builder /app /app

WORKDIR /app

ENTRYPOINT [ "./entrypoint.sh" ]
