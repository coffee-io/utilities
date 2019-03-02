package main;

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "os"
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
    remote_cksum := artifactLastChecksum(artifact);

    // if they are different, they changed
    if local_cksum == remote_cksum {
        fmt.Println("not changed ( key", local_cksum, ")");
    } else {
        fmt.Println("changed ( local", local_cksum, "remote", remote_cksum, ")")
    }
}

func artifactLastChecksum(artifact string) string {
    type Artifact struct {
        Artifact       string    `json:"artifact"`
        SourceChecksum string    `json:"sourceChecksum,omitempty"`
    }

    sess := session.Must(session.NewSession(&aws.Config{}))
    svc := dynamodb.New(sess)

    artifactKey := Artifact { Artifact: artifact }
    key, err := dynamodbattribute.MarshalMap(artifactKey)
    if (err != nil) {
        panic(err)
    }

    input := &dynamodb.GetItemInput{
        Key: key,
        TableName: aws.String("Artifacts"),
    }
    result, err := svc.GetItem(input)
    if (err != nil) {
        panic(err)
    }

    art := Artifact{}
    err = dynamodbattribute.UnmarshalMap(result.Item, &art)
    if (err != nil) {
        panic(err)
    }

    return art.SourceChecksum
}

// vim:ts=4:sts=4:sw=4:expandtab
