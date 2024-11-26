package api

import (
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

type GetResponse struct {
	Puzzle Puzzle `json:"puzzle"`
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
