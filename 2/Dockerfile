FROM alpine:3.19

RUN apk add --no-cache python3 py3-pip libpq && rm -rf /var/cache/apk/*

COPY requirements.txt .
RUN pip3 install --no-cache-dir --break-system-packages -r requirements.txt

RUN adduser -S -D -H -u 10001 appuser

USER appuser
WORKDIR /app

CMD ["sh"]
