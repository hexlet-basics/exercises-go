FROM golang:1.13

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN apt-get update && apt-get install -yqq git python3-pip

RUN pip3 install yamllint

WORKDIR /exercises-go
COPY . /exercises-go
