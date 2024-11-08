package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type RequestBody struct {
	Name string `json:"name"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (Response, error) {
	var reqBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return Response{
			StatusCode: http.StatusBadRequest,
			Body:       `{"message": "Invalid request"}`,
		}, nil
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       `{"message": "Hello ` + reqBody.Name + `"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
