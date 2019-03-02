package main;

import (
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "time"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "ARTIFACT_NAME DIRECTORY");
        os.Exit(1);
    }
    artifact := os.Args[1];
    directory := os.Args[2];

    // get checksums
    local_cksum := CalculateDirectoryChecksum(directory);
    updateDynamoTable(artifact, local_cksum);
}

func updateDynamoTable(artifact string, cksum string) {
    type Artifact struct {
        Artifact       string    `json:"artifact"`
        SourceChecksum string    `json:"sourceChecksum"`
        LastUpdated    time.Time `json:"lastUpdated"`
    }
    artifact_json := Artifact { artifact, cksum, time.Now() };

    sess := session.Must(session.NewSession(&aws.Config{}))
    svc := dynamodb.New(sess)

    av, err := dynamodbattribute.MarshalMap(artifact_json);
    if (err != nil) {
        panic(err)
    }

    input := &dynamodb.PutItemInput{
        Item: av,
        TableName: aws.String("Artifacts"),
    }

    _, err = svc.PutItem(input);
    if (err != nil) {
        panic(err)
    }

    fmt.Println(cksum)
}

// vim:ts=4:sts=4:sw=4:expandtab
