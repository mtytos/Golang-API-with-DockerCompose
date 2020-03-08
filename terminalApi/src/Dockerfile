FROM golang:1.13-alpine

ADD . /go/src/application
WORKDIR /go/src/application

RUN go build -o /application -mod=vendor

CMD ["/application"]
