FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o search .

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=builder --chown=appuser:appuser /app/search /app/search

USER appuser

ENTRYPOINT [ "/app/search" ]
