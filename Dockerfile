FROM golang

ADD . /go/src/github.com/rissicay/article

RUN go get github.com/tools/godep

RUN cd /go/src/github.com/rissicay/article && /go/bin/godep restore

RUN go install github.com/rissicay/article

ENTRYPOINT /go/bin/article

EXPOSE 8080
