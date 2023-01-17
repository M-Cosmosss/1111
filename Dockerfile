FROM ubuntu:latest

COPY ./test /test

ENTRYPOINT /test