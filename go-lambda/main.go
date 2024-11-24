package main

import (
	"encoding/json"
	"fmt"
	"goaws/events"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for /greet/{name}
func greetHandler(request events.LBRequest) (events.LBResponse, error) {
	name := request.PathParameters["name"]
	if name == "" {
		name = "Stranger"
	}
	message := fmt.Sprintf("Hello, %s!", name)

	// Prepare the response body
	response := map[string]string{
		"message": message,
	}

	// Convert to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return events.LBResponse{StatusCode: 500}, err
	}

	// Construct the LBResponse
	return events.LBResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseJSON),
	}, nil
}

// Handler for /health
func healthCheckHandler(request events.LBRequest) (events.LBResponse, error) {
	response := map[string]string{
		"status": "ok",
	}

	// Convert to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return events.LBResponse{StatusCode: 500}, err
	}

	// Construct the LBResponse
	return events.LBResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseJSON),
	}, nil
}

// Lambda function handler
func HandleRequest(request events.LBRequest) (events.LBResponse, error) {
	switch request.Path {
	case "/greet":
		return greetHandler(request)
	case "/health":
		return healthCheckHandler(request)
	default:
		return events.LBResponse{StatusCode: 404, Body: "Not Found"}, nil
	}
}

func main() {
	// Start the Lambda function
	lambda.Start(HandleRequest)
}
