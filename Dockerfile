FROM golang:1.10-alpine3.7

WORKDIR /go/src/github.com/Khigashiguchi/go-ecs-example/
COPY . /go/src/github.com/Khigashiguchi/go-ecs-example/

RUN apk update \
    && apk add --no-cache git

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go build -v -o server

EXPOSE 80
CMD ["./server"]
