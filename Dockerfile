# 使用官方的 Golang 镜像创建构建产物。
FROM golang:1.22 AS builder

RUN mkdir /app

WORKDIR /app
COPY . /app

RUN go env
RUN go env -w CGO_ENABLED=0
RUN go mod download

RUN go build -o cmd/ ./main

# --- runner ---
FROM debian as runner

RUN apt update && apt upgrade -y && apt install -y ca-certificates curl && update-ca-certificates


COPY --from=builder /app/cmd/main /usr/local/bin/main

RUN chmod +x /usr/local/bin/main

EXPOSE 8088

CMD ["/usr/local/bin/main"]
