package lambda_handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (lambdaHandler *Handler) HandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	contentType := map[string]string{
		"Content-Type": "application/json",
	}

	//Init Log
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	// Test Variable Environment
	logger.Info("Variable env ", os.Getenv("TEST_ENV"), nil)

	responseData, _ := json.Marshal(User{
		Name: "Sam",
	})

	return events.APIGatewayProxyResponse{
		Headers:    contentType,
		StatusCode: 200,
		Body:       string(responseData),
	}, nil
}
