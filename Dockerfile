FROM golang:1.19-alpine3.16 AS builder

WORKDIR /usr/local/go/src/

ADD . /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app main.go

FROM alpine:3.16

COPY --from=builder /usr/local/go/src/app /

CMD ["/app"]