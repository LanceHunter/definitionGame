package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// OxfordReply is the struct holding the data from the Oxford API's reply
type OxfordReply struct {
	Metadata OxfordMetadata `json:"metadata"`
	Results  []Result       `json:"results"`
}

// OxfordMetadata holds the metadata information from the API reply.
type OxfordMetadata struct {
	Operation string `json:"operation"`
	Provider  string `json:"provider"`
	Schema    string `json:"schema"`
}

// Result hold the result information from the API call.
type Result struct {
	ID             string          `json:"id"`
	Language       string          `json:"language"`
	Type           string          `json:"type"`
	Word           string          `json:"word"`
	LexicalEntries []LexicalEntry  `json:"lexicalEntries"`
	Pronunciations []Pronunciation `json:"pronunciations"`
}

// LexicalEntry is the struct with the lexical information from the reply.
type LexicalEntry struct {
	Language            string          `json:"language"`
	Text                string          `json:"text"`
	Compounds           []CompDer       `json:"compounds"`
	DerivativeOf        []CompDer       `json:"derivativeOf"`
	Derivatives         []CompDer       `json:"derivatives"`
	Entries             []Entry         `json:"entries"`
	GrammaticalFeatures []IDTextType    `json:"grammaticalFeatures"`
	LexicalCategory     IDText          `json:"lexicalCategory"`
	Notes               []IDTextType    `json:"notes"`
	PhrasalVerbs        []CompDer       `json:"phrasalVerbs"`
	Phrases             []CompDer       `json:"phrases"`
	Pronunciations      []Pronunciation `json:"pronunciations"`
	VariantForms        []VariantForm   `json:"variantForms"`
}

// Entry is the entry information for the lexical entry.
type Entry struct {
	HomographNumber       string          `json:"homographNumber"`
	CrossReferenceMarkers []string        `json:"crossReferenceMarkers"`
	CrossReference        []IDTextType    `json:"crossReferences"`
	Etymologies           []string        `json:"etymologies"`
	GrammaticalFeatures   []IDTextType    `json:"grammaticalFeatures"`
	Notes                 []IDTextType    `json:"notes"`
	Pronunciations        []Pronunciation `json:"pronunciations"`
	Senses                []Sense         `json:"senses"`
	Inflections           []Inflection    `json:"inflections"`
	VariantForms          []VariantForm   `json:"variantForms"`
}

// Sense is the struct for holding information on the senses of the word
type Sense struct {
	ID                    string          `json:"id"`
	Antonyms              []CompDer       `json:"antonyms"`
	Constructions         []Construction  `json:"constructions"`
	CrossReferenceMarkers []string        `json:"crossReferenceMarkers"`
	CrossReferences       []IDTextType    `json:"crossReferences"`
	Definitions           []string        `json:"definitions"`
	Domains               []IDText        `json:"domains"`
	Etymologies           []string        `json:"etymologies"`
	Examples              []Example       `json:"examples"`
	Inflections           []Inflection    `json:"inflections"`
	Notes                 []IDTextType    `json:"notes"`
	Pronunciations        []Pronunciation `json:"pronunciations"`
	Regions               []IDText        `json:"regions"`
	Registers             []IDText        `json:"registers"`
	ShortDefinitions      []string        `json:"shortDefinitions"`
	//subsenses is listed, but there was nothing in the struct so leaving out.
	Synonyms       []CompDer       `json:"synonyms"`
	ThesaurusLinks []ThesaurusLink `json:"thesaurusLinks"`
	VariantForms   []VariantForm   `json:"variantForms"`
}

// ThesaurusLink is the link for thesaurus information.
type ThesaurusLink struct {
	EntryID string `json:"entry_id"`
	SenseID string `json:"sense_id"`
}

// Example is the struct holding example uses of a word.
type Example struct {
	Definitions []string     `json:"definitions"`
	Domains     []IDText     `json:"domains"`
	Notes       []IDTextType `json:"notes"`
	Regions     []IDText     `json:"regions"`
	Registers   []IDText     `json:"registers"`
	SenseIDs    []string     `json:"senseIds"`
	Text        string       `json:"text"`
}

// Inflection is the struct for inflections of a word.
type Inflection struct {
	Domains             []IDText        `json:"domains"`
	GrammaticalFeatures []IDTextType    `json:"grammaticalFeatures"`
	InflectedForm       string          `json:"inflectedForm"`
	LexicalCategory     []IDText        `json:"lexicalCategory"`
	Pronunciations      []Pronunciation `json:"pronunciations"`
	Regions             []IDText        `json:"regions"`
	Registers           []IDText        `json:"registers"`
}

// Pronunciation is the struct with the pronunciation information from the
// reply.
type Pronunciation struct {
	AudioFile        string   `json:"audioFile"`
	Dialects         []string `json:"dialects"`
	PhoneticNotation string   `json:"phoneticNotation"`
	PhoneticSpelling string   `json:"phoneticSpelling"`
	Regions          []IDText `json:"regions"`
	Registers        []IDText `json:"registers"`
}

// CompDer is the compound/derivative information under LexicalEntry
type CompDer struct {
	Domains   []IDText `json:"domains"`
	ID        string   `json:"id"`
	Language  string   `json:"language"`
	Regions   []IDText `json:"regions"`
	Registers []IDText `json:"registers"`
	Text      string   `json:"text"`
}

// Construction is the struct much like CompDer, but like Variantform it's
// different enough that it needs its own struct.
type Construction struct {
	Domains   []IDText     `json:"domains"`
	Examples  []string     `json:"examples"`
	Notes     []IDTextType `json:"notes"`
	Regions   []IDText     `json:"regions"`
	Registers []IDText     `json:"registers"`
	Text      string       `json:"text"`
}

// VariantForm is for variant forms under the Entry struct.
type VariantForm struct {
	Domains        []IDText        `json:"domains"`
	Notes          []IDTextType    `json:"notes"`
	Pronunciations []Pronunciation `json:"pronunciations"`
	Regions        []IDText        `json:"regions"`
	Registers      []IDText        `json:"registers"`
	Text           string          `json:"text"`
}

// IDText is struct for the areas that have only the fields ID and Text.
// Region, Register, Domain, etc.
type IDText struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// IDTextType is the reference for CrossReference and GrammaticalFeatures under
// entry
type IDTextType struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

// Oxford is the function that is passed the randomly-selected word and calls
// the dictionary API to get the definition, putting it into the definition
// struct.
func Oxford(word string) string {
	log.Printf("The word is %s \n", word)

	// Get the API key info from environment variables.
	appID, idExists := os.LookupEnv("oxfordAppID")
	apiKey, keyExists := os.LookupEnv("oxfordAPIKey")

	// Verify that we got the variables, exit if we don't.
	log.Printf("The App ID is %s\n", appID)
	log.Printf("The API key is %s\n", apiKey)
	if !idExists || !keyExists {
		log.Fatalln("ERROR - appID or apiKey null.")
		return ""
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

	// Put the response into the struct.
	r := new(OxfordReply)
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatalln("Error parsing response body to JSON - ", err)
	}
	// Try printing out the struct
	log.Printf("%+v\n", r)

	// Print out a log of the body.
	log.Println("=====")
	log.Println(string(body))

	return r.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
}
