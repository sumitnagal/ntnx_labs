# -----------------------------------------------------------------------------
# Stage...
FROM golang:1.12.1-alpine AS build
WORKDIR /go/src/github.com/gopherland/ntnx_labs/deploy/picker

ENV VERSION=0.0.1

RUN apk --no-cache add git

COPY go.mod go.sum ./
COPY cmd cmd
COPY internal internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/picker \
  -ldflags="-w -s -X main.version=$VERSION" github.com/gopherland/ntnx_labs/deploy/picker/cmd


# -----------------------------------------------------------------------------
# Image...
FROM alpine:3.8

RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/picker /bin/picker
COPY assets assets

EXPOSE 4500
ENTRYPOINT ["/bin/picker"]