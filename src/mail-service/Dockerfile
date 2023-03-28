FROM golang:1.18-alpine

ARG LIBRDKAFKA_VERSION=2.0.2-r0

RUN apk add --upgrade apk-tools && \
    apk add libressl --update-cache --repository http://nl.alpinelinux.org/alpine/edge/main && \
    apk add librdkafka=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add librdkafka-dev=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

RUN apk add git protoc make curl gcc libc-dev
WORKDIR /app/igbot-sendmail
COPY . .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    export PATH="$PATH:$(go env GOPATH)/bin" && \
    PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
    curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip && \
    unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local && \
    mv $HOME/.local/include/google/ /usr/include/ && \
    make proto

ENV GO111MODULE=off
RUN go get github.com/githubnemo/CompileDaemon

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1
#ENV GOFLAGS -mod=vendor
ENV GOOS=linux
ENV GOARCH=arm64
ENV GO111MODULE=auto

RUN go mod download

ENTRYPOINT CompileDaemon -build="go build -tags musl -o sendmail_service ./cmd/server/" -command="./sendmail_service"
#ENTRYPOINT go run main.go