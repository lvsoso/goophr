FROM golang:1.16


ADD . /home/goophr/librarian

WORKDIR /home/goophr/librarian

RUN go build && cp librarian /go/bin/librarian

ENTRYPOINT /go/bin/librarian

EXPOSE 9000

