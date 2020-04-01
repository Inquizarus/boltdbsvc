FROM golang:alpine as builder

RUN apk update && apk add ca-certificates && apk add make

COPY . $GOPATH/src/github.com/inquizarus/golbag/

WORKDIR $GOPATH/src/github.com/inquizarus/golbag
RUN CGO111MODULE=on make build-linux
RUN mv golbag_unix /go/bin/golbag

FROM busybox:latest
COPY --from=builder /go/bin/golbag /go/bin/golbag
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
WORKDIR /go/bin
CMD ./golbag