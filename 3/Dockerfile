# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /src

COPY app/go.mod ./app/go.mod
COPY app/go.sum* ./app/
WORKDIR /src/app
RUN go mod download

COPY app/ ./
RUN go build -o /out/app .


# App stage
FROM alpine:3.20 AS app
RUN addgroup -S app && adduser -S -G app -h /app app

WORKDIR /app

COPY --from=build /out/app /bin/app

USER app

ENV APP_HOST=0.0.0.0 \
    APP_PORT=8080 \
    APP_READ_HEADER_TIMEOUT_SEC=5 \
    APP_SHUTDOWN_TIMEOUT_SEC=10

EXPOSE 8080

ENTRYPOINT ["/bin/app"]
