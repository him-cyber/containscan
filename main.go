// prototype complete — paused for re-evaluation
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/himaneesh/containscan/lambda_handler"
	"github.com/himaneesh/containscan/logging"
)

func main() {
	// If running inside AWS Lambda, trigger Lambda function
	if isLambda() {
		lambda.Start(lambda_handler.Handler)
		return
	}

	// Default local execution (for testing & development)
	logging.Info("ContainScan started successfully!")
	logging.Warning("Potential issue detected")
	logging.Error(nil)
	logging.Error(fmt.Errorf("critical failure occurred"))
}

// Detect if running inside AWS Lambda
func isLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}
