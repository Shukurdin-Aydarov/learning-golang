FROM golang

ADD . /go/src/learning-golang/ws-d
WORKDIR /go/src/learning-golang/ws-d

RUN go get github.com/lib/pq
RUN go install learning-golang/ws-d

ENTRYPOINT /go/bin/ws-d

EXPOSE 8080