package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func processBrainyLogRead(commandMap map[string]string) {
	searchText, containsSearchText := commandMap["l"]
	lineNumber, containsLineNumber := commandMap["n"]
	lineUUID, containsLineUUID := commandMap["u"]
	linesToShow := commandMap["m"]

	matchers := 0
	if containsSearchText {
		matchers++
	}
	if containsLineUUID {
		matchers++
	}
	if containsLineNumber {
		matchers++
	}

	if matchers != 1 {
		fmt.Println("Need exactly one of l, n, u to get!")
		return
	}

	searchType, isTask := commandMap["t"]
	if !isTask {
		searchType = "all"
	}

	_, hideMetadata := commandMap["nm"]
	if containsSearchText {
		getBrainyLogMatches(searchType, searchText, hideMetadata)
		return
	}
	if containsLineUUID {
		getUUIDMatches(lineUUID, hideMetadata, linesToShow)
		return
	}
	if containsLineNumber {
		lineUUID, err := getUUIDFromTemporaryPositionalNumber(lineNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		getUUIDMatches(lineUUID, hideMetadata, linesToShow)
		return
	}

}

func getUUIDMatches(lineUUID string, hideMetadata bool, linesToShow string) {
	var lineCount int
	if linesToShow == "" {
		lineCount = 0
	} else {
		var numberParseError error
		lineCount, numberParseError = strconv.Atoi(linesToShow)
		if numberParseError != nil {
			fmt.Println("Invalid value " + linesToShow + " for m!")
			return
		}
	}

	filename := defaultFilePath
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// this would contain an empty string after the last "\n", so subtracting 1
	lines := strings.Split(string(input), "\n")
	if len(lines) == 0 {
		fmt.Println("Empty file!")
		return
	}
	lines = lines[:len(lines)-1]

	for index, line := range lines {
		if getUUID(line) == lineUUID {
			if linesToShow == "" {
				displayLine(line, hideMetadata, "")
			} else {
				var startIndex int
				var endIndex int

				if index-lineCount >= 0 {
					startIndex = index - lineCount
				} else {
					startIndex = 0
				}
				if index+lineCount < len(lines) {
					endIndex = index + lineCount
				} else {
					endIndex = len(lines) - 1
				}

				positionalMappingToUUID := make(map[string]string)
				pos := 0

				for _, subline := range lines[startIndex : endIndex+1] {
					positionalMappingToUUID[strconv.Itoa(pos)] = getUUID(subline)
					displayLine(subline, hideMetadata, strconv.Itoa(pos))
					pos++
				}
				createPositionalMappingFile(positionalMappingToUUID)
			}

			break
		}
	}
}

func displayLine(line string, hideMetadata bool, positionalNumber string) {
	if hideMetadata {
		fmt.Println(getLineContent(line))
		return
	}
	if positionalNumber != "" {
		fmt.Println(line + " [" + positionalNumber + "]")
		return
	}
	fmt.Println(line)
	return
}

func getLineContent(line string) string {
	contentStartIndex := strings.Index(line, ">") + 1
	return line[contentStartIndex:]
}

func lineMatches(line string, searchType string, searchText string) (lineMatches bool) {
	lineContent := getLineContent(line)
	if !strings.Contains(strings.ToLower(lineContent), strings.ToLower(searchText)) {
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

func getBrainyLogMatches(searchType string, searchText string, hideMetadata bool) {
	fmt.Println("Getting matches for searchtext: ", searchText)
	file, err := os.Open(defaultFilePath)
	if err != nil {
		fmt.Println("Error opening file for get!")
	}
	defer file.Close()

	keywords := strings.Split(searchText, " ")

	scanner := bufio.NewScanner(file)

	positionalMappingToUUID := make(map[string]string)
	currentPos := 0

	for scanner.Scan() {
		currentLine := scanner.Text()

		for _, keyword := range keywords {
			if lineMatches(currentLine, searchType, keyword) {
				positionalMappingToUUID[strconv.Itoa(currentPos)] = getUUID(currentLine)
				if hideMetadata {
					fmt.Println(getLineContent(currentLine))
				} else {
					fmt.Println(currentLine + " [" + strconv.Itoa(currentPos) + "]")
				}
				currentPos++
				break
			}
		}
	}

	createPositionalMappingFile(positionalMappingToUUID)

	// err = ioutil.WriteFile("log-metadata.bl", []byte(output), 0644)
}

func createPositionalMappingFile(positionalMappingToUUID map[string]string) {

	metadataFile, err := os.Create("log-mapping.bl")
	defer metadataFile.Close()
	if err != nil {
		panic(err)
	}
	encoder := gob.NewEncoder(metadataFile)
	encodeError := encoder.Encode(positionalMappingToUUID)
	if encodeError != nil {
		panic(err)
	}
}
