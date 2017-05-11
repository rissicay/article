FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app
RUN go get github.com/tools/godep
RUN godep restore
RUN go-wrapper install
RUN go-wrapper download

COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080

CMD ["/go/bin/app"]
ENTRYPOINT ["/entrypoint.sh"]
