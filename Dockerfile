FROM golang:1.9.2-alpine3.6

WORKDIR /go/src/github.com/2at2/telegacli

COPY . .

RUN apk add --no-cache ca-certificates make git \
    && make deps \
    && make ok \
    && make release

ENTRYPOINT ["bin/telegacli"]
