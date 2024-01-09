# Builder
FROM golang:1.21.4-alpine3.17 as builder

RUN apk add alpine-sdk

WORKDIR /build

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -tags musl -o mts-sm-api ./cmd/rest-server

# Application container
FROM alpine:3.17

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/mts-sm-api /app/mts-sm-api

CMD ["/app/mts-sm-api"]