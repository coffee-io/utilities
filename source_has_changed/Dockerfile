FROM gobuild

WORKDIR /go/src/app
COPY . .

CMD go build . && aws s3 cp source_has_changed s3://coffee-artifacts/
