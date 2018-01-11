# lipsumgo

lipsumgo is a testing microservice with HTTP. The microservice generates "lorem ipsum" sentences and
prints standard output periodically. Also listens a HTTP port. HTTP responses include
a "lorem ipsum" sentence with some request headers: RequestURI, RemoteAddr.

## Usage

```sh
lipsumgo -n 60 -a :12345
lipsumgo -interval 60 -addr :12345
```

These examples generate a sentence one in 60 seconds and listens 12345 port (*:12345) for HTTP service.
If interval is 0, it never prints sentence to stdout. 60 seconds and ":12345" address are the default values.

## How to Install

```sh
go get github.com/orkunkaraduman/lipsumgo
```

## With Docker

```sh
git clone git@github.com:orkunkaraduman/lipsumgo.git
cd lipsumgo
docker build -t lipsumgo .
docker run --rm -it -p 12345:12345 lipsumgo -n 60 -a :12345
```
