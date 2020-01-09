FROM golang:1.9-alpine AS build

WORKDIR /go/src/github.com/orkunkaraduman/lipsumgo
COPY . .

RUN go-wrapper download
RUN go-wrapper install

ENTRYPOINT ["go-wrapper", "run"]


FROM alpine:3.7

RUN apk upgrade --no-cache && apk add --no-cache \
    openssl ca-certificates

WORKDIR /app
COPY --from=build /go/bin/lipsumgo .

ENTRYPOINT [ "./lipsumgo" ]
