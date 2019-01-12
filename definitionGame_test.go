package main_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	main "github.com/LanceHunter/definitionGame"
	alexa "github.com/arienmalec/alexa-go"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// // Setting up the structs for a correct Alexa request
	// type Session struct {
	// 	New         bool
	// 	SessionID   string
	// 	Application struct {
	// 		ApplicationID string
	// 	}
	// 	Attributes map[string]interface{}
	// 	User       struct {
	// 		UserID      string
	// 		AccessToken string
	// 	}
	// }
	//
	// type Resolutions struct {
	// 	ResolutionPerAuthority []struct {
	// 		Values []struct {
	// 			Value struct {
	// 				Name string
	// 				Id   string
	// 			}
	// 		}
	// 	}
	// }
	//
	// type Context struct {
	// 	System struct {
	// 		APIAccessToken string
	// 		Device         struct {
	// 			DeviceID string
	// 		}
	// 		Application struct {
	// 			ApplicationID string
	// 		}
	// 	}
	// }
	//
	// type Slot struct {
	// 	Name        string
	// 	Value       string
	// 	Resolutions Resolutions
	// }
	//
	// type Intent struct {
	// 	Name  string
	// 	Slots map[string]Slot
	// }
	//
	// type ReqBody struct {
	// 	Type        string
	// 	RequestID   string
	// 	Timestamp   string
	// 	Locale      string
	// 	Intent      Intent
	// 	Reason      string
	// 	DialogState string
	// }
	//
	// type Request struct {
	// 	Version string
	// 	Session Session
	// 	Body    ReqBody
	// 	Context Context
	// }
	// // End of the the structs for Alexa request.

	// Setting up the variable to hold first test result.
	var req1 alexa.Request
	var req2 alexa.Request

	// Importing the first JSON file
	jsonFile1, err := os.Open("testJSON/test1.json")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile1)
	json.Unmarshal([]byte(byteValue), &req1)

	// Importing the second JSON file
	jsonFile2, err := os.Open("testJSON/test2.json")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ = ioutil.ReadAll(jsonFile2)
	json.Unmarshal([]byte(byteValue), &req2)

	defer jsonFile1.Close()
	defer jsonFile2.Close()

	response1 := alexa.NewSimpleResponse("You did it!", "The test is working.")
	response2 := alexa.NewSimpleResponse("Error", "There was an error with this request.")

	tests := []struct {
		request alexa.Request
		expect  alexa.Response
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: req1,
			expect:  response1,
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: req2,
			expect:  response2,
			err:     main.ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.expect, response)
	}
}
