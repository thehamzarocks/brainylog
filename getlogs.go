package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processBrainyLogRead(commandMap map[string]string) {
	searchText := commandMap["l"]
	searchType, isTask := commandMap["t"]
	if !isTask {
		searchType = "all"
	}
	getBrainyLogMatches(searchType, searchText)
}

func lineMatches(line string, searchType string, searchText string) (lineMatches bool) {
	if !strings.Contains(strings.ToLower(line), strings.ToLower(searchText)) {
		return false
	}
	switch searchType {
	case "all":
		lineMatches = true
	case "allTasks":
		lineMatches = containsMetadata(line, "T")
	case "create":
		lineMatches = getMetadataValue(line, "T") == "0"
	case "progress":
		lineMatches = getMetadataValue(line, "T") == "1"
	case "suspend":
		lineMatches = getMetadataValue(line, "T") == "2"
	case "cancel":
		lineMatches = getMetadataValue(line, "T") == "3"
	case "complete":
		lineMatches = getMetadataValue(line, "T") == "4"
	default:
		fmt.Println("Invalid task type!")
		lineMatches = false
	}

	return lineMatches

}

func containsMetadata(line string, key string) bool {
	// Not a metadata line, ignore
	if string(line[0]) != "(" {
		return false
	}
	uuidEndIndex := strings.Index(line, ">")
	// this case should not happen
	if uuidEndIndex == -1 {
		return false
	}
	uuidStartIndex := strings.LastIndex(line[:uuidEndIndex], ")") + 1
	metadataStartIndex := strings.Index(line[:uuidStartIndex], "("+key+"-")
	return metadataStartIndex != -1
}

func getMetadataValue(line string, key string) string {
	// Not a metadata line, ignore
	if string(line[0]) != "(" {
		return ""
	}
	uuidEndIndex := strings.Index(line, ">")
	// this case should not happen
	if uuidEndIndex == -1 {
		return ""
	}
	uuidStartIndex := strings.LastIndex(line[:uuidEndIndex], ")") + 1
	metadataStartIndex := strings.Index(line[:uuidStartIndex], "("+key+"-")
	if metadataStartIndex == -1 {
		return ""
	}
	return string(line[metadataStartIndex+3])
}

func getBrainyLogMatches(searchType string, searchText string) {
	fmt.Println("Getting matches for searchtext: ", searchText)
	file, err := os.Open(defaultFilePath)
	if err != nil {
		fmt.Println("Error opening file for get!")
	}
	defer file.Close()

	keywords := strings.Split(searchText, " ")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()

		for _, keyword := range keywords {
			if lineMatches(currentLine, searchType, keyword) {
				fmt.Println(currentLine)
				break
			}
		}
	}
}
