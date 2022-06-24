FROM node:16.10.0-alpine

RUN apk add --no-cache bash \
    && npm install -g yarn@1.22.17 --force

USER node

WORKDIR /home/node/app
