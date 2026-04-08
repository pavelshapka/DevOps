FROM alpine:3.20

# TODO: поправить пользователя и группу
RUN apk add nginx && adduser -S -D -H -u 10001 usr

EXPOSE 8080

USER usr

CMD ["nginx", "-g", "daemon off;"]
