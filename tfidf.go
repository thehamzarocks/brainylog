package main

import (
	"fmt"
	"strings"

	"github.com/numbleroot/go-tfidf"
)

func getTfScore(line string, keywords []string) float64 {
	keywords = append(keywords, "to")

	tokenizedLine := tfidf.TokenizeDocument(line)

	tfScore := 0.0
	for _, keyWord := range keywords {
		tfScore += tfidf.TermFrequency(keyWord, true, tokenizedLine, tfidf.TermWeightingLog)
	}

	return tfScore / float64(len(line))
}

func getTfIdfScore() {
	s := "hello there, you must be hello"
	tokenizedS := tfidf.TokenizeDocument(s)
	normalizedFrequency := tfidf.TermFrequency("hello", true, tokenizedS, tfidf.TermWeightingLog)
	fmt.Println(normalizedFrequency)
}

func getExactMatchScore(line string, searchText string) float64 {
	exactMatchScore := 0.0
	if strings.Contains(line, searchText) {
		exactMatchScore += 2 * float64(len(searchText))
	}

	keyWords := strings.Split(searchText, " ")

	for _, keyword := range keyWords {
		if strings.Contains(line, keyword) {
			exactMatchScore += 1.0
		}
	}

	return exactMatchScore / float64(len(line))
}
