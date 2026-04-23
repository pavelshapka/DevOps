FROM alpine:3.20


RUN apk add nginx && adduser -S -D -H -u 10001 usr \
    && chown -R usr /var/lib/nginx /var/log/nginx /run/nginx /usr/share/nginx /etc/nginx

EXPOSE 8080

USER usr

CMD ["nginx", "-g", "daemon off;"]
