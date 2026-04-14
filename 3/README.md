Запуск:

```bash
docker compose up --build
```

Проверка:

```bash
curl -i http://localhost:8080/ping

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Tue, 14 Apr 2026 18:08:14 GMT
Content-Length: 5

pong
```

Повторный запуск не собирает приложение заново:

```bash
docker compose up --build

...
 => CACHED [app 2/4] RUN addgroup -S app && adduser -S -G app -h /app app                     0.0s
 => CACHED [app 3/4] WORKDIR /app                                                             0.0s
 => CACHED [build 2/8] WORKDIR /src                                                           0.0s
 => CACHED [build 3/8] COPY app/go.mod ./app/go.mod                                           0.0s
 => CACHED [build 4/8] COPY app/go.sum* ./app/                                                0.0s
 => CACHED [build 5/8] WORKDIR /src/app                                                       0.0s
 => CACHED [build 6/8] RUN go mod download                                                    0.0s
 => CACHED [build 7/8] COPY app/ ./                                                           0.0s
 => CACHED [build 8/8] RUN go build -o /out/app .                                             0.0s
 => CACHED [app 4/4] COPY --from=build /out/app /bin/app                                      0.0s
...
```
