# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image)
FROM golang:1-buster AS builder
RUN go version

COPY . /go/src/github.com/amnk/dd2tf/
WORKDIR /go/src/github.com/amnk/dd2tf/
RUN set -x && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -v
RUN set -x && \
    go get -u github.com/go-bindata/go-bindata/... && \
    go generate -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o dd2tf .

# Stage 2 (to create a downsized "container executable")

FROM alpine:3
RUN apk --no-cache add ca-certificates && mkdir -p /app/exports
WORKDIR /app/
COPY --from=builder /go/src/github.com/amnk/dd2tf/dd2tf /app/dd2tf

ENTRYPOINT ["/app/dd2tf"]
