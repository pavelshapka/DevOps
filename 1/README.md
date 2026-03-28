# Docker: запуск nginx

## Сборка образа

```bash
docker build -t mynginx .
```

## Запуск

```bash
docker-compose up
```

## Проверка доступности

```bash
curl http://localhost:8080
```
