FROM gobuild

WORKDIR /go/src/app
COPY . .

CMD go build -v -ldflags="-s -w" -o update_source_cksum . && aws s3 cp ./update_source_cksum s3://coffee-artifacts/
