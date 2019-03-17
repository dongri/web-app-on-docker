FROM golang:1.12-alpine
LABEL maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --no-cache git
RUN go get -u github.com/githubnemo/CompileDaemon

ADD . /go/src/github.com/dongri/web-app-on-docker/go
WORKDIR /go/src/github.com/dongri/web-app-on-docker/go

EXPOSE 3001

CMD PORT=3001 CompileDaemon -command="./go"
