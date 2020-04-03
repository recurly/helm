# Buildtime image
FROM golang:1.14-alpine3.11 AS builder
RUN apk add --no-cache make bash git

WORKDIR /go/src/k8s.io/helm
COPY . ./
RUN make bootstrap build

# Runtime image
FROM alpine:3.11

RUN apk add --no-cache ca-certificates socat

ENV HOME /tmp

COPY --from=builder /go/src/k8s.io/helm/bin/tiller /tiller

EXPOSE 44134
USER 65534
ENTRYPOINT ["/tiller"]