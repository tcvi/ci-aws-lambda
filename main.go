package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type MyEvent struct {
	Name string `json:"name"`
}

func hello(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	name := os.Getenv("NAME")
	return fmt.Sprintf("Hello name: %s", name), nil
}

func main() {
	lambda.Start(hello)
}
