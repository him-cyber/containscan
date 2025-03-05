package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event struct for receiving container launch events
type Event struct {
	Detail struct {
		ClusterArn string `json:"clusterArn"`
		TaskArn    string `json:"taskArn"`
	} `json:"detail"`
}

// AWS Lambda function handler
func Handler(ctx context.Context, event Event) (string, error) {
	log.Printf("Received container event: %+v\n", event)

	cluster := event.Detail.ClusterArn
	task := event.Detail.TaskArn
	log.Printf("Scanning container from Cluster: %s, Task: %s", cluster, task)

	// TODO: Add container security scanning logic here

	return fmt.Sprintf("ContainScan executed successfully for Task: %s", task), nil
}

func main() {
	lambda.Start(Handler)
}
