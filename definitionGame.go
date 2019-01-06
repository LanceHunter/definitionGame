package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Response struct, with string for message
type Response struct {
	Message string
}

// Handler for lambda
func Handler() (Response, error) {
	return Response{Message: "This is a test."}, nil
}

// Main function
func main() {
	lambda.Start(Handler)
}
