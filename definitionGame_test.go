package main_test

import (
	"testing"

	main "github.com/LanceHunter/definitionGame"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "Lance"},
			expect:  "Hello Lance",
			err:     nil,
		},
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "%%"},
			expect:  "Hello %%",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     main.ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		/*
			response, err := main.Handler(test.request)
			assert.IsType(t, test.err, err)
			assert.Equal(t, test.expect, response.Body)
		*/
	}
}
