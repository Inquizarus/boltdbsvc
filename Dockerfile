FROM golang:alpine as builder

RUN apk update && apk add ca-certificates && apk add make

COPY . $GOPATH/src/github.com/inquizarus/boltdbsvc/

WORKDIR $GOPATH/src/github.com/inquizarus/boltdbsvc
RUN CGO111MODULE=on make build-linux
RUN mv boltdbsvc_unix /go/bin/boltdbsvc

FROM busybox:latest
COPY --from=builder /go/bin/boltdbsvc /go/bin/boltdbsvc
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
WORKDIR /go/bin
CMD ./boltdbsvc