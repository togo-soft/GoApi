FROM golang:latest

LABEL maintainer="xuthus5@gmail.com" version=1.0

WORKDIR /Project

RUN go env -w GOPROXY=https://goproxy.cn,direct

EXPOSE 80

ENTRYPOINT ["go","run","main.go"]
