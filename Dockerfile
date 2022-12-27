FROM golang:1.18-alpine AS base

RUN apk upgrade --no-cache && apk add --no-cache \
    make git

FROM base AS builder

RUN apk upgrade --no-cache && apk add --no-cache \
    make git

WORKDIR /src
COPY . .
RUN make vendor
RUN make build
RUN mv target/ /app/
RUN cp -a docker/entrypoint.sh /app/

FROM alpine:3.16

RUN apk upgrade --no-cache && apk add --no-cache \
    bash openssl ca-certificates

COPY --from=builder /app /app

WORKDIR /app

ENTRYPOINT [ "./entrypoint.sh" ]
