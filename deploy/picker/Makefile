PACKAGE      := github.com/gopherland/ntnx_labs/deploy

SVC_VERSION  := 0.0.1
SVC_IMAGE    := picker
SVC_APP      := picker
SVC_PACKAGE  := ${PACKAGE}/${SVC_APP}
SVC_IMG_NAME := ${SVC_IMAGE}:${SVC_VERSION}

default: help

build:  ## Build local picker
	@go build -v -o execs/picker \
  -ldflags="-w -s -X main.version=${SVC_VERSION}" ${SVC_PACKAGE}/cmd

img:  ## Build Picker Svc Docker Image
	@docker build --rm -t ${SVC_IMG_NAME} .

run:  ## Run dockerized Picker
	@docker run -it -p 4500:4500 ${SVC_IMG_NAME} -d

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;168m %s\n", $$1, $$2}'
