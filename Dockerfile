FROM golang

RUN go get github.com/gorilla/mux

RUN go get gopkg.in/intercom/intercom-go.v2

WORKDIR /go/src/github.com/oms-services/intercom

ADD . /go/src/github.com/oms-services/intercom

RUN go install github.com/oms-services/intercom

ENTRYPOINT intercom

EXPOSE 3000