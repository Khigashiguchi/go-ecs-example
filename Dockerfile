FROM golang:1.10-alpine3.7

WORKDIR /go/src/github.com/Khigashiguchi/go-ecs-example/
COPY . /go/src/github.com/Khigashiguchi/go-ecs-example/

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh \
    bash \
    curl \
    python

# Install the AWS CLI
# https://aws.amazon.com/jp/blogs/news/managing-secrets-for-amazon-ecs-applications-using-parameter-store-and-iam-roles-for-tasks/
RUN curl -O https://bootstrap.pypa.io/get-pip.py
RUN python get-pip.py
RUN pip install awscli

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go build -v -o server

EXPOSE 80
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./server"]
