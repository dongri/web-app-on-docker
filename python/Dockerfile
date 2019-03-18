FROM python:3.7.2-alpine3.9

RUN pip3 install flask

ADD . /src
WORKDIR /src

EXPOSE 3000

CMD python3 run.py
