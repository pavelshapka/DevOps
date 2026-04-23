FROM node:25-alpine AS build

WORKDIR /app

COPY ./package.json ./package-lock.json* ./

RUN npm ci && npm cache clean --force

COPY . .

RUN npm run build && rm -rf node_modules && npm cache clean --force

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY ./frontend.nginx.conf /etc/nginx/nginx.conf
