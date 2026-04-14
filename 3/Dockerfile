# Build stage
FROM alpine:3.23 AS build

RUN apk add --no-cache go=1.25.9-r0 && rm -rf /var/cache/apk/*

WORKDIR /app
COPY app/go.mod ./go.mod
COPY app/go.sum* ./
RUN go mod download

COPY app/ ./
RUN go build -o /app/bin .

# App stage
FROM alpine:3.23

RUN mkdir -p /app
COPY --from=build /app/bin /app/bin

RUN adduser -S -D -H -u 10001 appuser
USER appuser

ENV APP_HOST=0.0.0.0 \
    APP_PORT=8080 \
    APP_READ_HEADER_TIMEOUT_SEC=5 \
    APP_SHUTDOWN_TIMEOUT_SEC=10

EXPOSE 8080

ENTRYPOINT ["/app/bin"]
