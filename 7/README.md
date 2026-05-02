# Решение находится в репозитории на Gitlab

[Приложение с CI/CD пайплайном](https://gitlab.com/shapka.pa/devops-cicd-app.git)

Изначально:

```bash
~ curl -v "http://93.77.190.239:1709/calc?op=add&a=5&b=15"
*   Trying 93.77.190.239:1709...
* Connected to 93.77.190.239 (93.77.190.239) port 1709
> GET /calc?op=add&a=5&b=15 HTTP/1.1
> Host: 93.77.190.239:1709
> User-Agent: curl/8.7.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Sat, 02 May 2026 14:52:07 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 93.77.190.239 left intact
20%
```

Приложение обновилось:

```bash
➜  ~ curl -v "http://93.77.190.239:1709/calc?op=mod&a=20&b=15"
*   Trying 93.77.190.239:1709...
* Connected to 93.77.190.239 (93.77.190.239) port 1709
> GET /calc?op=mod&a=20&b=15 HTTP/1.1
> Host: 93.77.190.239:1709
> User-Agent: curl/8.7.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Sat, 02 May 2026 15:00:18 GMT
< Content-Length: 1
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 93.77.190.239 left intact
5%
```
