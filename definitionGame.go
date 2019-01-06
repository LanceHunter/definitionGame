package main

import (
	"errors"
	"log"
	"math/rand"

	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("You done messed up.")
)

// Handler for lambda
func Handler(request alexa.Request) (alexa.Response, error) {

	// The array of words that we'll be using for the game.
	words := [5]string{"parcel", "calculator", "trail", "gold", "push"}

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request \n")

	randomWord := words[rand.Intn(5)]
	log.Printf("The random word is %s\n", randomWord)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body.Type) < 1 {
		log.Printf("request.QueryStringParameters is %+v\n", request.Body)
		log.Printf("------\n")

		return alexa.NewSimpleResponse("Error", "There was an error"), ErrNameNotProvided
	}

	log.Printf("request.Body is %+v\n", request.Body)

	return alexa.NewSimpleResponse("Help for Hello", "To receive a greeting, ask hello to say hello"), nil

}

// Main function
func main() {
	lambda.Start(Handler)
}
