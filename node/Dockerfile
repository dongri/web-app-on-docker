FROM node:8.15.1-alpine
LABEL maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --update alpine-sdk

ENV APP_ROOT /src
COPY . $APP_ROOT
WORKDIR $APP_ROOT
RUN npm install -g nodemon
RUN npm install

EXPOSE 3000
CMD nodemon .bin/www
