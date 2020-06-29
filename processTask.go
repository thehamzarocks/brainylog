package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func processTask(commandMap map[string]string) {
	taskUUID, containsTaskUUID := commandMap["u"]
	temporaryPositionalNumber, containstemporaryPositionalNumber := commandMap["n"]

	if !containsTaskUUID && !containstemporaryPositionalNumber {
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

	if containsTaskUUID {
		changeTaskState(taskUUID, toState)
		return
	}
	if containstemporaryPositionalNumber {
		matchingUUID := getUUIDFromTemporaryPositionalNumber(temporaryPositionalNumber)
		changeTaskState(matchingUUID, toState)
		return
	}
}

func getUUIDFromTemporaryPositionalNumber(temporaryPositionalNumber string) string {
	decodeFile, err := os.Open("log-mapping.bl")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	positionalMappings := make(map[string]string)
	decoder.Decode(&positionalMappings)

	uuid, containsUUID := positionalMappings[temporaryPositionalNumber]
	if !containsUUID {
		panic("Unable to find line matching positional number " + temporaryPositionalNumber + "!")
	}
	return uuid
}

func changeTaskState(taskUUID string, toState string) {
	filename := defaultFilePath
	// fmt.Println(filename, line)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if getUUID(line) == taskUUID {
			changedLine := changeTaskStateAtLine(line, toState)
			if changedLine == "" {
				panic("Error occured while changing task state!")
			}
			lines[i] = changedLine
			fmt.Println("Task state changed for line:")
			fmt.Println(changedLine)
			break
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(defaultFilePath, []byte(output), 0644)

	if err != nil {
		fmt.Println("Something went wrong while editing task state!")
	}
}

func getUUID(line string) string {
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
	metadataStartIndex := strings.Index(line[:uuidStartIndex], "(T-")
	return line[:metadataStartIndex+3] + value + line[metadataStartIndex+4:]
}
