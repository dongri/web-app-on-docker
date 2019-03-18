FROM ruby:2.5.1-alpine3.7
LABEL maintainer "Dongri Jin <dongrify@gmail.com>"

RUN apk add --update alpine-sdk

ENV APP_ROOT /src

COPY . $APP_ROOT

WORKDIR $APP_ROOT

RUN bundle install

EXPOSE 3000

CMD bundle exec rackup -p 3000 -o 0.0.0.0
