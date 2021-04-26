FROM golang:1.16


ADD . /go/src/github.com/last-ent/distributed-go/chapter5/goophr/concierge

WORKDIR /go/src/github.com/last-ent/distributed-go/chapter5/goophr/concierge

RUN go install github.com/last-ent/distributed-go/chapter5/goophr/concierge

ENTRYPOINT /go/bin/concierge

EXPOSE 9000

