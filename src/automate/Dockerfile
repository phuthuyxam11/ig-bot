FROM golang:1.18-alpine

RUN apk add git
WORKDIR /app/igbot-automate
#COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=off

RUN go get github.com/githubnemo/CompileDaemon
#RUN go get -t -d github.com/tebeka/selenium
ENTRYPOINT CompileDaemon -build="go build" -command="./igbot-automate"
#ENTRYPOINT go run main.go