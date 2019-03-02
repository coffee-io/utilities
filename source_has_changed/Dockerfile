FROM gobuild

WORKDIR /go/src/app
COPY . .

CMD go build -v -ldflags="-s -w" -o source_has_changed . && aws s3 cp ./source_has_changed s3://coffee-artifacts/
