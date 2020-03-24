FROM golang:1.13 AS builder

# Go module proxy
RUN go env -w GOPROXY=https://goproxy.cn,direct

# Compile the project
WORKDIR ${GOPATH}/src/github.com/xiaogaozi/tikv-proxy
COPY . .
RUN make server

FROM debian:10

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /go/src/github.com/xiaogaozi/tikv-proxy/_build/tikv-proxy /app/tikv-proxy
ENTRYPOINT ["./tikv-proxy"]
