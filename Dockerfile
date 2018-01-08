FROM golang:1.9-alpine

WORKDIR /go/src/github.com/orkunkaraduman/lipsumgo
COPY . .

RUN go-wrapper download
RUN go-wrapper install

ENTRYPOINT ["go-wrapper", "run"]
