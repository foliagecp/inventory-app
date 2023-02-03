FROM golang:1.19 as builder

LABEL maintainer="NJWS, Inc."

WORKDIR /src/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/inventory ./cmd/inventory

FROM ubuntu:18.04

LABEL maintainer="NJWS, Inc."

RUN apt update && \
    apt install ca-certificates netcat -y && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/inventory /usr/bin/

RUN chmod +x /usr/bin/inventory

CMD ["/usr/bin/inventory" ,"agent", "run"]
