package main

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("you done messed up, son")
)

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) (alexa.Response, error) {
	var response alexa.Response
	var err error
	switch request.Body.Intent.Name {
	case "hello":
		response, err = handleHello(request)
	default:
		response = alexa.NewSimpleResponse("Test", "You're in a weird place right now, dude.")
	}
	return response, err
}

func handleHello(request alexa.Request) (alexa.Response, error) {

	wordNumber := rand.Intn(5)
	randomWord := words[wordNumber]
	log.Printf("The random number is %d\n", wordNumber)
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

// Handler for lambda
func Handler(request alexa.Request) (alexa.Response, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request \n")

	return DispatchIntents(request)
}

// Function that is testing the call to the Oxford API.
func Oxford(word string) {
	log.Printf("The word is %s \n", word)
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln("Error with GET call - ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body - ", err)
	}

	log.Println(string(body))

}

// Main function
func main() {
	Oxford("test")
	lambda.Start(Handler)
}
