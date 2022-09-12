FROM golang:1.19.0-alpine3.16 as builder

RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 \
    go build  -ldflags="-w -s" -tags musl -o ./exec ./exec.go

FROM alpine:3.16 as production

RUN apk update && apk add --no-cache curl ca-certificates
RUN rm -rf /var/cache/apk/*

COPY --from=builder --chown=65534:0 /app/exec /go/bin/exec

USER 65534

WORKDIR /go/bin

ADD docker-entrypoint.sh docker-entrypoint.sh

ENTRYPOINT ["./docker-entrypoint.sh"]