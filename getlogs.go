package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type lineProcessorFn func(lines []string, argsMap map[string]string) (writeBack bool)

func processFile(processLines lineProcessorFn, argsMap map[string]string) {

	filename := defaultFilePath
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to find log! Check if log.bl exists in the current directory and it's not currently open!")
		return
	}

	// this would contain an empty string after the last "\n", so subtracting 1
	lines := strings.Split(string(input), "\n")
	if len(lines) == 0 {
		fmt.Println("Empty file!")
		return
	}
	lines = lines[:len(lines)-1]

	writeBack := processLines(lines, argsMap)

	if writeBack {
		output := strings.Join(lines, "\n")
		output += "\n"
		err = ioutil.WriteFile(defaultFilePath, []byte(output), 0644)

		if err != nil {
			fmt.Println("Something went wrong while writing to file!")
		}
	}
}

func getMatchScore(line string, searchType string, keywords []string) (score float64) {
	searchText := strings.Join(keywords, " ")
	if !lineMatches(line, searchType, searchText) {
		return -1
	}

	lineContent := getLineContent(line)
	tfScore := getTfScore(lineContent, keywords)
	exactMatchScore := getExactMatchScore(lineContent, searchText)
	return tfScore + exactMatchScore
}

func isDeleted(line string) bool {
	return getMetadataValue(line, "S", 2) == "01"
}

func getSearchTextMatches(lines []string, argsMap map[string]string) (writeBack bool) {
	searchText := argsMap["searchText"]
	searchType := argsMap["searchType"]
	hideMetadata := argsMap["hideMetadata"]

	fmt.Println("Getting matches for searchtext: ", searchText+"\n")

	matchingItems := make(map[string]float64, 0)

	for _, line := range lines {
		if isDeleted(line) {
			continue
		}

		keywords := strings.Split(searchText, " ")

		matchScore := getMatchScore(line, searchType, keywords)

		if matchScore > 0 {
			matchingItems[line] = matchScore
		}
	}

	scoreRankedMatches := getPriorityRankedItems(matchingItems)
	// for _, keyword := range keywords {
	// 	if lineMatches(line, searchType, keyword) {
	// 		matchingItems[line] = getMatchScore(line)
	// 		positionalMappingToUUID[strconv.Itoa(currentPos)] = getUUID(line)
	// 		if hideMetadata == "hideMetadata" {
	// 			fmt.Println(getLineContent(line))
	// 		} else {
	// 			fmt.Println(getLineContent(line) + " [" + strconv.Itoa(currentPos) + "]")
	// 		}
	// 		currentPos++
	// 		break
	// 	}
	// }

	positionalMappingToUUID := make(map[string]string)
	currentPos := 0

	for _, match := range scoreRankedMatches {
		positionalMappingToUUID[strconv.Itoa(currentPos)] = getUUID(match)
		if hideMetadata == "hideMetadata" {
			fmt.Println(getLineContent(match))
		} else {
			fmt.Println(getLineContent(match) + " [" + strconv.Itoa(currentPos) + "]")
		}
		currentPos++
	}

	fmt.Println("")
	createPositionalMappingFile(positionalMappingToUUID)

	return false
}

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

	searchType, isTask := commandMap["t"]
	if !isTask {
		searchType = "all"
	}

	_, hideMetadata := commandMap["nm"]

	var argsMap map[string]string = make(map[string]string)

	if hideMetadata {
		argsMap["hideMetadata"] = "hideMetadata"
	} else {
		argsMap["hideMetadata"] = "showMetadata"
	}

	if matchers == 0 {
		argsMap["searchText"] = ""
		argsMap["searchType"] = searchType
		processFile(getSearchTextMatches, argsMap)
	}

	if containsSearchText {
		argsMap["searchText"] = searchText
		argsMap["searchType"] = searchType
		processFile(getSearchTextMatches, argsMap)
		return
	}
	if containsLineUUID {
		argsMap["uuid"] = lineUUID
		argsMap["linesToShow"] = linesToShow
		processFile(getUUIDMatches, argsMap)
		return
	}
	if containsLineNumber {
		lineUUID, err := getUUIDFromTemporaryPositionalNumber(lineNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		argsMap["uuid"] = lineUUID
		argsMap["linesToShow"] = linesToShow
		processFile(getUUIDMatches, argsMap)
		return
	}

}

func getUUIDMatches(lines []string, argsMap map[string]string) (writeBack bool) {
	lineUUID := argsMap["uuid"]
	linesToShow := argsMap["linesToShow"]
	hideMetadata := argsMap["hideMetadata"]

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

	fmt.Println("Getting matches: " + "\n")

	for index, line := range lines {
		if getUUID(line) == lineUUID {
			if linesToShow == "" {
				displayLine(line, hideMetadata, "")
			} else {
				displayLinesInRange(lines, index, lineCount, hideMetadata)
			}
			break
		}
	}
	fmt.Println("")
	return false
}

func displayLinesInRange(lines []string, index int, lineCount int, hideMetadata string) {
	var startIndex, endIndex int

	currentIndex := index - 1
	linesCounted := 0
	for currentIndex >= 0 && linesCounted < lineCount {
		if !isDeleted(lines[currentIndex]) {
			linesCounted++
		}
		currentIndex--
	}
	startIndex = currentIndex + 1

	currentIndex = index + 1
	linesCounted = 0
	for currentIndex < (len(lines)) && linesCounted < lineCount {
		if !isDeleted(lines[currentIndex]) {
			linesCounted++
		}
		currentIndex++
	}
	endIndex = currentIndex

	// if index-lineCount >= 0 {
	// 	startIndex = index - lineCount
	// } else {
	// 	startIndex = 0
	// }
	// if index+lineCount < len(lines) {
	// 	endIndex = index + lineCount
	// } else {
	// 	endIndex = len(lines) - 1
	// }

	positionalMappingToUUID := make(map[string]string)
	pos := 0

	for _, subline := range lines[startIndex : endIndex+1] {
		if isDeleted(subline) {
			continue
		}
		positionalMappingToUUID[strconv.Itoa(pos)] = getUUID(subline)
		displayLine(subline, hideMetadata, strconv.Itoa(pos))
		pos++
	}

	// create new positional mappings only if line metadata is displayed
	if hideMetadata != "hideMetadata" {
		createPositionalMappingFile(positionalMappingToUUID)
	}
}

func displayLine(line string, hideMetadata string, positionalNumber string) {
	if hideMetadata == "hideMetadata" {
		if getMetadataValue(line, "S", 2) == "01" {
			fmt.Println("This line has been deleted: " + line)
			return
		}
		fmt.Println(getLineContent(line))
		return
	}

	if positionalNumber != "" {
		fmt.Println(getLineContent(line) + " [" + positionalNumber + "]")
		return
	}

	if getMetadataValue(line, "S", 2) == "01" {
		fmt.Println("This line has been deleted: " + line)
		return
	}
	fmt.Println(getLineContent(line))
	return
}

func getLineContent(line string) string {
	contentStartIndex := strings.Index(line, ">") + 1
	return line[contentStartIndex:]
}

// returns true/false based on whether the line matches any keyword in the searchText
func lineMatches(line string, searchType string, searchText string) (lineMatches bool) {
	lineContent := getLineContent(line)

	keyWords := strings.Split(searchText, " ")
	containsMatchingKeyword := false

	for _, keyWord := range keyWords {
		if strings.Contains(strings.ToLower(lineContent), strings.ToLower(keyWord)) {
			containsMatchingKeyword = true
			break
		}
	}

	if !containsMatchingKeyword {
		return false
	}

	switch searchType {
	case "all":
		lineMatches = true
	case "allTasks":
		lineMatches = containsMetadata(line, "T")
	case "create":
		lineMatches = getMetadataValue(line, "T", 1) == "0"
	case "progress":
		lineMatches = getMetadataValue(line, "T", 1) == "1"
	case "suspend":
		lineMatches = getMetadataValue(line, "T", 1) == "2"
	case "cancel":
		lineMatches = getMetadataValue(line, "T", 1) == "3"
	case "complete":
		lineMatches = getMetadataValue(line, "T", 1) == "4"
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

func getMetadataValue(line string, key string, valueLength int) string {
	// Not a metadata line, ignore
	if len(line) == 0 {
		return ""
	}
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
	return string(line[metadataStartIndex+2+len(key) : metadataStartIndex+2+len(key)+valueLength])
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
