Запуск:

```bash
docker compose build system && docker compose build build && docker compose up app
```

Проверка:

```bash
curl -i http://localhost:8080/ping

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Tue, 14 Apr 2026 21:05:16 GMT
Content-Length: 5

pong
```
