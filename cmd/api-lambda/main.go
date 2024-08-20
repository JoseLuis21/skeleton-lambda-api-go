package main

import (
	lambda_handler "github.com/JoseLuis21/skeleton-lambda-api-go/internal/lambda-handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handlerLambda := lambda_handler.NewHandler()
	lambda.Start(handlerLambda.HandlerRequest)
}
