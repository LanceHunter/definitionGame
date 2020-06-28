package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// OxfordReply is the struct holding the data from the Oxford API's reply
type OxfordReply struct {
	Metadata OxfordMetadata `json:"metadata"`
	Results  []Results      `json:"results"`
}

// OxfordMetadata holds the metadata information from the API reply.
type OxfordMetadata struct {
	Operation string `json:"operation"`
	Provider  string `json:"provider"`
	Schema    string `json:"schema"`
}

// Results hold the result information from the API call.
type Results struct {
	ID             string           `json:"id"`
	Language       string           `json:"language"`
	Type           string           `json:"type"`
	Word           string           `json:"word"`
	LexicalEntries []LexicalEntry   `json:"lexicalEntries"`
	Pronunciations []Pronunciations `json:"pronunciations"`
}

// LexicalEntry is the struct with the lexical information from the reply.
type LexicalEntry struct {
	Word string
}

// Pronunciations is the struct with the pronunciation information from the
// reply.
type Pronunciations struct {
}

// Oxford is the function that is passed the randomly-selected word and calls
// the dictionary API to get the definition, putting it into the definition
// struct.
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
