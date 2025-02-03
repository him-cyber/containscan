package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

// 🔹 Define logger at the package level (so all functions can access it)
var logger *log.Logger

// 🔹 AWS DynamoDB client
var dynamoClient *dynamodb.Client

func init() {
	// 🔹 Initialize logger BEFORE anything else
	file, err := os.OpenFile("containscan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("[ERROR] Failed to open log file, defaulting to console output")
		logger = log.New(os.Stdout, "", 0)
	} else {
		logger = log.New(file, "", 0)
	}

	// 🔹 Setup AWS SDK config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		logger.Fatal("Unable to load AWS SDK config: ", err)
	}

	// 🔹 Create DynamoDB client
	dynamoClient = dynamodb.NewFromConfig(cfg)
}

// 🔹 Function to log structured JSON messages
func logJSON(level, message string) {
	entry := map[string]string{
		"level":     level,
		"timestamp": time.Now().Format(time.RFC3339),
		"message":   message,
	}

	jsonLog, _ := json.Marshal(entry)
	logger.Println(string(jsonLog)) // ✅ Now `logger` is recognized

	putItemToDynamoDB(entry)
}

// 🔹 Send logs to DynamoDB
func putItemToDynamoDB(entry map[string]string) {
	_, err := dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ContainScanLogs"),
		Item: map[string]types.AttributeValue{
			"log_id":    &types.AttributeValueMemberS{Value: uuid.New().String()},
			"level":     &types.AttributeValueMemberS{Value: entry["level"]},
			"timestamp": &types.AttributeValueMemberS{Value: entry["timestamp"]},
			"message":   &types.AttributeValueMemberS{Value: entry["message"]},
		},
	})
	if err != nil {
		fmt.Println("[ERROR] Failed to send log to DynamoDB:", err)
	}
}

// 🔹 Public logging functions
func Info(message string) {
	logJSON("INFO", message)
}

func Warning(message string) {
	logJSON("WARNING", message)
}

func Error(err error) {
	if err != nil {
		logJSON("ERROR", err.Error())
	}
}
