FROM golang:1.11.10-alpine3.9 AS build
RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /go/src

COPY vendor/ ./

WORKDIR /go/src/ClickPhonebook

COPY db db
COPY handler handler
COPY logger logger
COPY schema schema
COPY util util

COPY main.go main.go

RUN go install ./...

FROM alpine:3.9
RUN apk --no-cache add curl
EXPOSE 8080/tcp
COPY templates /templates
WORKDIR /usr/bin
COPY --from=build /go/bin .
