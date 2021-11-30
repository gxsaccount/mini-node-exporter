#build exporter
FROM golang:1.17 AS builder
WORKDIR /mini-node-exporter
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN make

#set application
FROM alpine:latest
COPY --from=builder /mini-node-exporter/bin/ /app/
WORKDIR /app

CMD ["./exporter"]
