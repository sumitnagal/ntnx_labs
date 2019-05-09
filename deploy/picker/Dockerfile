# -----------------------------------------------------------------------------
# Stage...
FROM golang:1.12.1-alpine AS build
WORKDIR /go/src/github.com/gopherland/ntnx_labs/deploy/picker

ENV VERSION=0.0.2 GO111MODULE=on

# RUN apk --no-cache add git

COPY go.mod go.sum main.go ./
COPY internal internal


RUN apk --no-cache add git ;\
  CGO_ENABLED=0 GOOS=linux go build -o /go/bin/picker \
  -ldflags="-w -s -X main.version=$VERSION" github.com/gopherland/ntnx_labs/deploy/picker


# -----------------------------------------------------------------------------
# Image...
FROM alpine:3.8

COPY --from=build /go/bin/picker /bin/picker
COPY assets assets

EXPOSE 4500
ENTRYPOINT ["/bin/picker"]