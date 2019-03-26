FROM golang:alpine as builder
ENV GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN adduser -D -g '' service
RUN apk add --no-cache git
ADD . /go/src/gapi
WORKDIR /go/src/gapi
RUN go build && cd other && go build -a -o /go/src/gapi/other/healthcheck
USER service

FROM alpine:3.7
WORKDIR /go/src/gapi
RUN apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Shanghai
COPY --from=builder /go/src/gapi .
ENTRYPOINT ["/go/src/gapi/gapi"]
HEALTHCHECK --start-period=2s --interval=10s --timeout=5s CMD ["/go/src/gapi/other/healthcheck"]