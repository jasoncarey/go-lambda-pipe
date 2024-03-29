package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/cobra"
)

var invokeLambdaCmd = &cobra.Command{
	Use:   "invokeLambda [function name] [input file]",
	Short: "Invoke a Lambda function with the contents of a JSON file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		invokeLambdaFunction(args[0], args[1])
	},
}

func invokeLambdaFunction(functionName, jsonFilePath string) {
	jsonFileContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-2"),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	lambdaClient := lambda.NewFromConfig(cfg)

	invokeOutput, err := lambdaClient.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      jsonFileContent,
	})
	if err != nil {
		log.Fatalf("Failed to invoke Lambda function: %v", err)
	}

	fmt.Printf("Lambda returned status code: %d\n", invokeOutput.StatusCode)
	if len(invokeOutput.Payload) > 0 {
		fmt.Println("Response:")
		fmt.Println(string(invokeOutput.Payload))
	}
}

func init() {
	rootCmd.AddCommand(invokeLambdaCmd)
}
