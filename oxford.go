package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Oxford that is testing the call to the Oxford API.
func Oxford(word string) {
	log.Printf("The word is %s \n", word)

	// Get the API key info from environment variables.
	appID, idExists := os.LookupEnv("oxfordAppID")
	apiKey, keyExists := os.LookupEnv("oxfordAPIKey")

	// Verify that we got the variables, exit if we don't.
	log.Printf("The App ID is %s\n", appID)
	log.Printf("The API key is %s\n", apiKey)
	if !idExists || !keyExists {
		log.Fatalln("ERROR - appID or apiKey null.")
		return
	}

	client := &http.Client{}
	reqURL := "https://od-api.oxforddictionaries.com//api/v2/entries/en-us/" + word
	log.Println("Request URL - ", reqURL)
	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("app_id", appID)
	req.Header.Add("app_key", apiKey)

	// Make the call and get the response.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error with GET call - ", err)
	}

	// Close out response (deferred).
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body - ", err)
	}

	// Print out a log of the body.
	log.Println(string(body))
}
