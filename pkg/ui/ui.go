package ui

import (
	"fmt"
	"math/rand"
	"spelling_bee/pkg/api"
	"strings"
	"time"
)

func convertToUpper(l []string) []string {
	convertedStrings := []string{}
	for _, s := range l {
		convertedStrings = append(convertedStrings, strings.ToUpper(s))
	}

	return convertedStrings
}

func shuffleLetters(l []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(l), func(i, j int) { l[i], l[j] = l[j], l[i] })
	return l
}

func RenderGet(puzzle api.Puzzle) string {
	yellow := "\033[33m"
	reset := "\033[0m"

	otherLetters := shuffleLetters(convertToUpper(puzzle.OtherLetters))
	requiredLetter := fmt.Sprint(yellow + strings.ToUpper(puzzle.RequiredLetter) + reset)

	letters := append(otherLetters[:3], append([]string{requiredLetter}, otherLetters[3:]...)...)

	return strings.Join(letters, " ")
}

func RenderGuess(score api.Score) string {
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	switch score.Response {
	case "word":
		return fmt.Sprintf(yellow + "Valid word!" + reset)
	case "panagram":
		return fmt.Sprintf(green + "Panagram!" + reset)
	case "invalid":
		return fmt.Sprint(red + "Not a valid word..." + reset)
	default:
		return score.Response
	}
}
