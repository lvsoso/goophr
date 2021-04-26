FROM golang:1.16


ADD . /home/goophr/concierge

WORKDIR /home/goophr/concierge

RUN go build && cp concierge /go/bin/concierge

ENTRYPOINT /go/bin/concierge

EXPOSE 9000

