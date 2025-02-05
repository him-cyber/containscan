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
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/google/uuid"
)

// Logger setup
var logger *log.Logger
var snsClient *sns.Client
var snsTopicArn string
var awsEnabled bool

func init() {
	// Open local log file
	file, err := os.OpenFile("containscan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("[ERROR] Failed to open log file, using console output")
		logger = log.New(os.Stdout, "", 0)
	} else {
		logger = log.New(file, "", 0)
	}

	// Get AWS region from environment (default to us-east-1)
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1" // Default region
	}

	// Check AWS credentials
	awsAccessKey, awsKeyExists := os.LookupEnv("AWS_ACCESS_KEY_ID")
	awsSecretKey, awsSecretExists := os.LookupEnv("AWS_SECRET_ACCESS_KEY")

	if awsKeyExists && awsSecretExists {
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
			config.WithRegion(region),
		)

		if err != nil {
			fmt.Println("[WARNING] AWS detected, but failed to initialize SDK. Running in local mode.")
			awsEnabled = false
			return
		}

		// Initialize SNS Client
		snsClient = sns.NewFromConfig(cfg)
		snsTopicArn = "arn:aws:sns:" + region + ":160885291126:ContainScanAlerts" // Uses dynamic region
		awsEnabled = true
		fmt.Println("[INFO] AWS SNS alerts enabled in region:", region)
	} else {
		fmt.Println("[INFO] AWS not configured. Running in local mode.")
		awsEnabled = false
	}
}

// Log function with optional SNS alerting
func logJSON(level, message string) {
	entry := map[string]string{
		"log_id":     uuid.New().String(),
		"level":      level,
		"utc_time":   time.Now().UTC().Format(time.RFC3339),
		"local_time": time.Now().Local().Format(time.RFC3339),
		"message":    message,
	}

	jsonLog, _ := json.Marshal(entry)
	logger.Println(string(jsonLog)) // Write to local log file

	// Send alert for ERROR logs
	if level == "ERROR" && awsEnabled {
		sendSNSAlert(message)
	}
}

// Send alert via AWS SNS
func sendSNSAlert(message string) {
	_, err := snsClient.Publish(context.TODO(), &sns.PublishInput{
		Message:  aws.String("ContainScan ALERT: " + message),
		TopicArn: aws.String(snsTopicArn),
	})

	if err != nil {
		fmt.Println("[ERROR] Failed to send SNS alert:", err)
	}
}

// Public log functions
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
