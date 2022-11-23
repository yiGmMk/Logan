FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
RUN export GOPATH=/build && \
    cd src/logan && \
    go mod vendor && \
    CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o /build/logan ./server/cmd

FROM ubuntu

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/logan /app/ServerGo

CMD ["./ServerGo","-c","/config/logan.toml"]
