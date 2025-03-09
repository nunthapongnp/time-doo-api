FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN if [ -f go.sum ]; then \
      cp go.sum .; \
    else \
      touch go.sum; \
    fi

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /time-doo-api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /time-doo-api .

EXPOSE 3000

ENTRYPOINT ["/time-doo-api"]
