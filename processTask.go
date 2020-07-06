package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func processTask(commandMap map[string]string) {
	taskUUID, containsTaskUUID := commandMap["u"]
	temporaryPositionalNumber, containsTemporaryPositionalNumber := commandMap["n"]

	if !containsTaskUUID && !containsTemporaryPositionalNumber {
		fmt.Println("Task processing needs a task UUID or temporary positional number!")
		return
	}

	toState, containsToState := commandMap["t"]
	if !containsToState {
		fmt.Println("Task processing needs a task to-state!")
		return
	}
	if toState != "create" && toState != "progress" && toState != "suspend" && toState != "cancel" && toState != "complete" {
		fmt.Println("Invalid value " + toState + " for task toState!")
	}

	argsMap := make(map[string]string)

	argsMap["toState"] = toState

	if containsTaskUUID {
		argsMap["taskUUID"] = taskUUID
		processFile(changeTaskState, argsMap)
		return
	}
	if containsTemporaryPositionalNumber {
		matchingUUID, err := getUUIDFromTemporaryPositionalNumber(temporaryPositionalNumber)
		if err != nil {
			fmt.Println(err)
		}
		argsMap["taskUUID"] = matchingUUID
		processFile(changeTaskState, argsMap)
		return
	}
}

func getUUIDFromTemporaryPositionalNumber(temporaryPositionalNumber string) (string, error) {
	decodeFile, err := os.Open("log-mapping.bl")
	if err != nil {
		return "", errors.New("Unable to find mappings file")
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	positionalMappings := make(map[string]string)
	decoder.Decode(&positionalMappings)

	uuid, containsUUID := positionalMappings[temporaryPositionalNumber]
	if !containsUUID {
		return "", errors.New("Unable to find line matching positional number " + temporaryPositionalNumber + "!")
	}
	return uuid, nil
}

func changeTaskState(lines []string, argsMap map[string]string) (linesToWrite []string, shouldWriteLines bool) {
	taskUUID := argsMap["taskUUID"]
	toState := argsMap["toState"]

	for i, line := range lines {
		if getUUID(line) == taskUUID {
			if isDeleted(line) {
				fmt.Println("Line has been deleted: " + line)
				return
			}
			changedLine := changeTaskStateAtLine(line, toState)
			if changedLine == "" {
				panic("Error occured while changing task state!")
			}
			lines[i] = changedLine
			fmt.Println("Task state changed for line:")
			fmt.Println("")
			fmt.Println(changedLine)
			fmt.Println("")
			break
		}
	}

	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(defaultFilePath, []byte(output), 0644)

	if err != nil {
		fmt.Println("Something went wrong while editing task state!")
	}

	return lines, true
}

func getUUID(line string) string {
	if len(line) == 0 {
		return ""
	}

	// Not a metadata line, ignore
	if string(line[0]) != "(" {
		return ""
	}
	uuidEndIndex := strings.Index(line, ">")
	// this case should not happen
	if uuidEndIndex == -1 {
		fmt.Println("wow no")
		return ""
	}
	uuidStartIndex := strings.LastIndex(line[:uuidEndIndex], ")") + 1
	return line[uuidStartIndex:uuidEndIndex]
}

func changeTaskStateAtLine(line string, toState string) string {
	toStateValue := "9"
	switch toState {
	case "create":
		toStateValue = "0"
	case "progress":
		toStateValue = "1"
	case "suspend":
		toStateValue = "2"
	case "cancel":
		toStateValue = "3"
	case "complete":
		toStateValue = "4"
	}

	return setMetadataValue(line, "T", toStateValue)
}

func setMetadataValue(line string, key string, value string) string {
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
	return line[:metadataStartIndex+len(key)+2] + value + line[metadataStartIndex+2+len(key)+len(value):]
}
