FROM node:18-alpine AS build

WORKDIR /client

COPY package.json .
COPY yarn.lock .
COPY .env .

RUN yarn install

COPY . .

RUN yarn build

FROM nginx:1.25.3-alpine

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*

COPY --from=build /client/build .
COPY nginx.conf /etc/nginx/
COPY cert.pem key.unencrypted.pem /

ENTRYPOINT [ "nginx", "-g", "daemon off;" ]
