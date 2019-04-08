FROM haskell:8.6.3
LABEL maintainer "Dongri Jin <dongrify@gmail.com>"

ENV APP_ROOT /src
ADD . $APP_ROOT
WORKDIR $APP_ROOT/web

RUN stack setup
RUN stack build

EXPOSE 3000
CMD stack exec web-exe
