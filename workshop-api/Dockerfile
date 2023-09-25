FROM golang:1.20.8-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download \
    ; CGO_ENABLED=0 go build -o bin/api cmd/api/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/api .

ENTRYPOINT /app/api