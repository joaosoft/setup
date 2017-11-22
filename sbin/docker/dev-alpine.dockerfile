FROM golang:1.9-alpine

ARG PROJECT_NAME=go-mock

# install git and mercurial
RUN apk add --update git mercurial && rm -rf /var/cache/apk/*

# install make
RUN apk add --update bash make && rm -rf /var/cache/apk/*

# install dep
RUN go get -u github.com/golang/dep/cmd/dep

# install dependencies
ADD Gopkg.toml Gopkg.lock /go/src/$PROJECT_NAME/
RUN cd /go/src/$PROJECT_NAME && dep ensure -vendor-only

# copy configuration
ADD ./config /etc/$PROJECT_NAME

# add source code
ADD . /go/src/$PROJECT_NAME/
WORKDIR /go/src/$PROJECT_NAME/

EXPOSE 8080
ENTRYPOINT ["go"]
