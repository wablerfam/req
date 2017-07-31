FROM golang:1.8.3-alpine3.6

RUN set -ex \
  && apk add --no-cache \
    curl \
    git

COPY ./patch/os_linux.go /usr/local/go/src/runtime/os_linux.go

RUN set -ex \
  && go get github.com/golang/dep/cmd/dep \
  && go get github.com/laher/goxc
