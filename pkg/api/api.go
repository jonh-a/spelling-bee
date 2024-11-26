package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Puzzle struct {
	Date           string   `json:"date"`
	RequiredLetter string   `json:"requiredLetter"`
	OtherLetters   []string `json:"otherLetters"`
	PanagramCount  int64    `json:"panagramCount"`
	ValidWordCount int64    `json:"validWordCount"`
}

type Score struct {
	Date           string `json:"date"`
	Response       string `json:"response"`
	PanagramCount  int64  `json:"panagramCount"`
	ValidWordCount int64  `json:"validWordCount"`
}

type GetResponse struct {
	Puzzle Puzzle `json:"puzzle"`
}

type GuessRequest struct {
	Word string `json:"word"`
}

func getBaseUrl() string {
	return "http://localhost:3000/spelling"
}

func GetPuzzle(date string) (Puzzle, error) {
	response, err := http.Get(fmt.Sprintf("%s?date=%s", getBaseUrl(), date))
	if err != nil {
		return Puzzle{}, fmt.Errorf("error fetching puzzle: %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Puzzle{}, fmt.Errorf("error reading puzzle: %w", err)
	}

	var getResponse GetResponse
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return Puzzle{}, fmt.Errorf("error unmarsalling puzzle: %w", err)
	}

	return getResponse.Puzzle, nil
}

func GuessPuzzle(date string, guess string) (Score, error) {
	request := GuessRequest{guess}
	requestJson, err := json.Marshal(request)
	if err != nil {
		return Score{}, fmt.Errorf("error marshalling json: %w", err)
	}

	response, err := http.Post(fmt.Sprintf("%s/guess?date=%s", getBaseUrl(), date), "application/json", bytes.NewBuffer(requestJson))
	if err != nil {
		return Score{}, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Score{}, fmt.Errorf("error reading response: %w", err)
	}

	var score Score
	err = json.Unmarshal(body, &score)
	if err != nil {
		return Score{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return score, nil
}
