FROM golang

RUN go get github.com/gorilla/mux

RUN go get gopkg.in/intercom/intercom-go.v2

WORKDIR /go/src/github.com/heaptracetechnology/microservice-intercom

ADD . /go/src/github.com/heaptracetechnology/microservice-intercom

RUN go install github.com/heaptracetechnology/microservice-intercom

ENTRYPOINT microservice-intercom

EXPOSE 3000