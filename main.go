package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	s "strings"
	"time"
)

func main() {
	results := make(chan map[string]int)
	wordCounts := make(map[string]int)

	for i := range 10 {
		go worker(i, results)
	}

	for i := 0; i < 10; i++ {
		wordCount := <-results
		mergeIntoWithAdd(wordCounts, wordCount)
	}

	topWords := sortedByValue(wordCounts)

	fmt.Printf("Top 10 words by count: %s\n", topWords[len(topWords)-10:])
}

func worker(id int, results chan<- map[string]int) {
	delay := rand.IntN(10)
	fmt.Printf("Worker %d will delay for %ds\n", id, delay)
	time.Sleep(time.Second * time.Duration(delay))

	lipsum := randomLipsum()
	wordCounts := wordCount(lipsum)
	results <- wordCounts
}

func mergeIntoWithAdd(into, m map[string]int) {
	for k, v := range m {
		into[k] += v
	}
}

func randomLipsum() string {
	resp, err := http.Get(fmt.Sprintf("https://loripsum.net/api/%d/long/plaintext", rand.IntN(10)+1))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func strip(str string) string {
	var stripped s.Builder
	for _, r := range str {
		if isAlpha(r) {
			stripped.WriteRune(r)
		}
	}
	return stripped.String()
}

func clean(str string) string {
	return s.ToLower(strip(str))
}

func wordCount(str string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range s.Fields(str) {
		cleaned := clean(word)
		wordCount[cleaned]++
	}
	return wordCount
}
