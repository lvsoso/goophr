FROM golang:1.16

ADD . /home/goophr/littlefs

WORKDIR /home/goophr/littlefs

RUN go build && cp littlefs /go/bin/littlefs

ENTRYPOINT /go/bin/littlefs

