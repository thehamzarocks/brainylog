package main

import (
	"bufio"
	"os"
	"errors"
	"strings"
	"fmt"
)

func processGetArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "hello", errors.New("Please provide a search text for get")
	}
	searchText := strings.Join(args[1:], " ")
	return searchText, nil
}

func getBrainyLogMatches(searchText string) {
	fmt.Println("Getting matches for searchtext: ", searchText)
	file, err := os.Open("bin/log.bl")
	if (err != nil) {
		fmt.Println("Error opening file for get!")
	}
	defer file.Close()

	keywords := strings.Split(searchText, " ")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()

		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(currentLine), strings.ToLower(keyword)) {
				fmt.Println(currentLine)
				break
			}
		}
	}
}