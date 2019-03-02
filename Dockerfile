FROM golang:1.12

RUN apt-get update \
	&& apt-get install -y awscli git \
	&& go get github.com/aws/aws-sdk-go/aws \
	&& go get github.com/aws/aws-sdk-go/aws/session \
	&& go get github.com/aws/aws-sdk-go/service/dynamodb \
	&& go get github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute
