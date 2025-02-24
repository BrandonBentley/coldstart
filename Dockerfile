FROM golang:1.24-alpine3.20 AS builder

WORKDIR /

RUN apk update && apk upgrade
RUN apk add make

COPY . .
RUN GOARCH=amd64 go build -o coldstart .

# Deploy container
FROM alpine:3.20 AS deploy

WORKDIR /

COPY --from=builder coldstart /bin/coldstart
COPY --from=builder config/env /config/env

RUN apk update && apk upgrade
RUN rm -rf /var/cache/apk/*

ENTRYPOINT [ "coldstart" ]
