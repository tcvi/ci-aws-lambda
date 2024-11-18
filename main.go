package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type Response struct {
	Name string `json:"name"`
}

func hello(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	name := os.Getenv("NAME")
	return Response{
		fmt.Sprintf("Hello %s", name),
	}, nil
}

func main() {
	lambda.Start(hello)
}
