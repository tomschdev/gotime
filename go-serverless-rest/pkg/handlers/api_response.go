package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	stringBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal("error during json marshal of body")
	}
	resp.Body = string(stringBody)
	return &resp, nil
}
