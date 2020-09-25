include make.conf
# Variables from make.conf:
#
# DOCKER_REPO

SHELL := /bin/bash
APPNAME := postal-rest

VERSION := $(shell cat VERSION | tr -d "[:space:]")
COMMIT := $(shell git rev-parse --short HEAD)

default: compile

clean:
	rm -Rf src
	rm -Rf bin

compile:
	export GOPATH=$(shell pwd); go build -o bin/${APPNAME}

docker:
	docker build . -t ${DOCKER_REPO}/${APPNAME}:${VERSION}-${COMMIT}
	docker push ${DOCKER_REPO}/${APPNAME}:${VERSION}-${COMMIT}
