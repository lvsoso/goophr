FROM golang:1.16


ADD . /go/src/github.com/last-ent/distributed-go/chapter5/goophr/librarian

RUN go install github.com/last-ent/distributed-go/chapter5/goophr/librarian

ENTRYPOINT /go/bin/librarian

EXPOSE 9000

