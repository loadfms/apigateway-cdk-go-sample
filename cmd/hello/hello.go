package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (Response, error) {
	return Response{
		StatusCode: http.StatusOK,
		Body:       `{"message": "Hello route healthy"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
