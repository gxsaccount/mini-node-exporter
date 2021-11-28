#build exporter
FROM golang:1.16.7 AS dev
WORKDIR /mini-node-exporter
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN make

#set application
FROM alpine:latest
COPY --from=dev /mini-node-exporter/bin/ /app/
WORKDIR /app

CMD ["./exporter"]