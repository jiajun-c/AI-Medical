FROM golang:alpine

LABEL maintainer="chengjiajun20@gmail.com"
COPY . /ai
WORKDIR /ai
VOLUME /pictures
RUN go env -w GO111MODULE=on&&go env -w GOPROXY=https://goproxy.cn,direct
CMD go run ./app