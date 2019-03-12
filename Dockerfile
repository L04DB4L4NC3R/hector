FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/hector

ADD . /go/src/github.com/angadsharma1016/hector
WORKDIR /go/src/github.com/angadsharma1016/hector
RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
EXPOSE 3000
ENTRYPOINT watcher -run github.com/angadsharma1016/hector/ -watch github.com/angadsharma1016/hector