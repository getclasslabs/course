FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat
COPY ./course /go/src/
COPY ./config.yaml /

CMD ["/go/src/course}"]