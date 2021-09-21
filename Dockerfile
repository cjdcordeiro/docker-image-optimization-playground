FROM circleci/golang:latest

USER root

COPY . /opt

WORKDIR /opt/code

RUN apt-get update

# debugging tools
RUN apt-get install -y vim  

# app dependencies
RUN apt-get install -y libusb-1.0 libusb-dev udev

RUN go mod tidy && go build


