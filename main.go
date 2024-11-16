package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func hello(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	return "Hello APIGatewayProxyRequest", nil
}

func main() {
	lambda.Start(hello)
}
