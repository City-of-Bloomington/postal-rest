FROM ubuntu:latest
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    golang autoconf automake build-essential curl git libsnappy-dev libtool pkg-config

RUN git clone https://github.com/openvenues/libpostal /usr/local/src/libpostal
WORKDIR /usr/local/src/libpostal
RUN /bin/bash -c './bootstrap.sh; ./configure --datadir=/srv/data/libpostal; make; make install; ldconfig'

WORKDIR /srv/sites/postal-rest
ENV GOPATH=/srv/sites/postal-rest
ENV LISTEN_PORT=80
COPY main.go main.go
RUN go get .; go build -o bin/postal-rest

EXPOSE 80
ENTRYPOINT bin/postal-rest
