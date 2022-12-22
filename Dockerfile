FROM golang:1.19.2 AS builder
WORKDIR /build-dir
COPY go.mod .
COPY go.sum .
RUN go mod download all
COPY api api
COPY cmd cmd
COPY config config
RUN go build -o /tmp/liteapi github.com/kosrk/ton-http-liteapi/cmd/liteapi

FROM ubuntu:20.04 AS ton-http-liteapi
RUN apt-get update
COPY --from=builder /tmp/liteapi /app/liteapi
CMD ["/app/liteapi", "-v"]