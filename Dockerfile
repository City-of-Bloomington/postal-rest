FROM ubuntu:latest
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    golang autoconf automake build-essential curl git libsnappy-dev libtool pkg-config

RUN git clone https://github.com/openvenues/libpostal /usr/local/src/libpostal
WORKDIR /usr/local/src/libpostal
RUN /bin/bash -c './bootstrap.sh; ./configure --datadir=/srv/data/libpostal; make; make install; ldconfig'

WORKDIR /srv/sites/postal-rest
ENV GOPATH=/srv/sites/postal-rest
RUN go get github.com/City-of-Bloomington/postal-rest; go install github.com/City-of-Bloomington/postal-rest

EXPOSE 8080
ENTRYPOINT bin/postal-rest
