FROM golang:1.21.0 AS builder
ARG VERSION=v0.0.1

WORKDIR /src

COPY ../../../go.mod ../../../go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
COPY ../../../.. /src
RUN mkdir -p bin/ && go build -ldflags "-X main.Version=$VERSION" -o ./bin/ ./app/prom_agent/...

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/app/prom_agent/configs /data/conf

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
VOLUME /app/log
VOLUME /app/cache

# TZ=Asia/Shanghai
ENV TZ=Asia/Shanghai

CMD ["./prom_agent", "-conf", "/data/conf"]