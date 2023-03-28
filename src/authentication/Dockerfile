FROM golang:1.18-alpine

ARG LIBRDKAFKA_VERSION=2.0.2-r0

RUN apk add --upgrade apk-tools && \
    apk add libressl --update-cache --repository http://nl.alpinelinux.org/alpine/edge/main && \
    apk add librdkafka=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add librdkafka-dev=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

RUN apk add git gcc libc-dev
WORKDIR /app/igbot-auth
#COPY . .
ENV GO111MODULE=off
RUN go get github.com/githubnemo/CompileDaemon

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1
#ENV GOFLAGS -mod=vendor
ENV GOOS=linux
ENV GOARCH=arm64
ENV GO111MODULE=auto

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# RUN go mod tidy
#RUN go get github.com/githubnemo/CompileDaemon
#RUN go get -t -d github.com/tebeka/selenium
ENTRYPOINT CompileDaemon -build="go build -tags musl" -command="./authentication"
#ENTRYPOINT go run main.go