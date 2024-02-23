package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func apiWordAPI(partOfSpeech string) string {
	type WordData struct {
		// Replace with actual data structure based on your API response
		Word       string `json:"word"`
		Definition string `json:"definition"`
		// ... other fields
	}

	// Make the HTTP request
	url := "https://wordsapiv1.p.rapidapi.com/words/?partOfSpeech=" + partOfSpeech + "&frequencyMin=2&random=true&limit=1"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_RAPIDAPI"))
	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var data WordData
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}

	return data.Word
}
