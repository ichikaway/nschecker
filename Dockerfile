# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/nschecker .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/nschecker /usr/local/bin/nschecker

ENTRYPOINT ["nschecker"]
