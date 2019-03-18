FROM python:3.7.2-alpine3.9
LABEL maintainer "Dongri Jin <dongrify@gmail.com>"

RUN pip3 install flask

ADD . /src
WORKDIR /src

EXPOSE 3000

CMD python3 run.py
