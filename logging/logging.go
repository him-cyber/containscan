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

// Global logger
var logger *log.Logger

// AWS DynamoDB client (optional)
var dynamoClient *dynamodb.Client
var awsEnabled bool

func init() {
	// Open local log file
	file, err := os.OpenFile("containscan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("[ERROR] Failed to open log file, defaulting to console output")
		logger = log.New(os.Stdout, "", 0)
	} else {
		logger = log.New(file, "", 0)
	}

	// Check if AWS credentials exist
	_, awsKeyExists := os.LookupEnv("AWS_ACCESS_KEY_ID")
	_, awsSecretExists := os.LookupEnv("AWS_SECRET_ACCESS_KEY")

	if awsKeyExists && awsSecretExists {
		// Load AWS Config
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		if err != nil {
			fmt.Println("[WARNING] AWS credentials found but could not initialize AWS SDK. Running in local mode.")
			awsEnabled = false
			return
		}

		// Initialize DynamoDB Client
		dynamoClient = dynamodb.NewFromConfig(cfg)
		awsEnabled = true
		fmt.Println("[INFO] AWS logging enabled.")
	} else {
		fmt.Println("[INFO] No AWS credentials found. Running in local mode.")
		awsEnabled = false
	}
}

// Log function with optional AWS storage
func logJSON(level, message string) {
	entry := map[string]string{
		"level":      level,
		"utc_time":   time.Now().UTC().Format(time.RFC3339),
		"local_time": time.Now().Local().Format(time.RFC3339),
		"message":    message,
	}

	jsonLog, _ := json.Marshal(entry)
	logger.Println(string(jsonLog)) // Write to local log file

	// Store in AWS if enabled
	if awsEnabled {
		putItemToDynamoDB(entry)
	}
}

// Store logs in AWS DynamoDB
func putItemToDynamoDB(entry map[string]string) {
	_, err := dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ContainScanLogs"),
		Item: map[string]types.AttributeValue{
			"log_id":     &types.AttributeValueMemberS{Value: uuid.New().String()},
			"level":      &types.AttributeValueMemberS{Value: entry["level"]},
			"utc_time":   &types.AttributeValueMemberS{Value: entry["utc_time"]},
			"local_time": &types.AttributeValueMemberS{Value: entry["local_time"]},
			"message":    &types.AttributeValueMemberS{Value: entry["message"]},
		},
	})
	if err != nil {
		fmt.Println("[ERROR] Failed to send log to DynamoDB:", err)
	}
}

// Public logging functions
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
