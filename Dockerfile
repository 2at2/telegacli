FROM alpine:3.6

COPY bin/ /usr/bin

RUN set -ex && apk update && apk add ca-certificates

ENTRYPOINT ["telegacli"]
