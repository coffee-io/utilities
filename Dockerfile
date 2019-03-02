FROM golang:1.12-alpine

RUN apk add aws-cli --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ --allow-untrusted \
	&& apk add git \
	&& go get github.com/aws/aws-sdk-go/aws \
	&& go get github.com/aws/aws-sdk-go/aws/session \
	&& go get github.com/aws/aws-sdk-go/service/dynamodb \
	&& go get github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute
