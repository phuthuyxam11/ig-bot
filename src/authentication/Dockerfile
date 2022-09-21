FROM golang:1.18-alpine

RUN apk add git
WORKDIR /app/igbot-auth
#COPY . .
ENV GO111MODULE=off
RUN go get github.com/githubnemo/CompileDaemon

ENV CGO_ENABLED=0
ENV GO111MODULE=auto

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# RUN go mod tidy
#RUN go get github.com/githubnemo/CompileDaemon
#RUN go get -t -d github.com/tebeka/selenium
ENTRYPOINT CompileDaemon -build="go build" -command="./authentication"
#ENTRYPOINT go run main.go