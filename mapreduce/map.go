package mapreduce

import (
	"strings"
)

// Map function splits text into words and counts them.
func Map(text string) map[string]int {
	wordCounts := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,!?;"))
		wordCounts[word]++
	}
	return wordCounts
}
