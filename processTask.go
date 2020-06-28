package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func processTask(args []string) {
	taskUUID, toState := processTaskArgs(args)
	if toState == "error" {
		fmt.Println("Invalid task parameters!")
		return
	}
	changeTaskState(taskUUID, toState)
}

func processTaskArgs(args []string) (string, string) {
	if len(args) < 3 {
		return "", "error"
	}

	taskUUID := args[1]
	toState := args[2]

	if (toState != "create" && toState != "progress" && toState != "suspend" && toState != "cancel" && toState != "complete") {
		return taskUUID, "error"
	}

	return taskUUID, toState
}

func changeTaskState(taskUUID string, toState string) {
	filename := "bin/log.bl"
	// fmt.Println(filename, line)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if getUUID(line) == taskUUID {
			changedLine := changeTaskStateAtLine(line, toState)
			if (changedLine == "") {
				panic("Error occured while changing task state!")
			}
			lines[i] = changedLine
			fmt.Println("Task state changed for line:")
			fmt.Println(changedLine)
			break
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("bin/log.bl", []byte(output), 0644)

	if err != nil {
		fmt.Println("Something went wrong while editing task state!")
	}
}

func getUUID(line string) (string) {
	// Not a metadata line, ignore
	if string(line[0]) != "(" {
		return ""
	}
	uuidEndIndex := strings.Index(line, ">");
	// this case should not happen
	if uuidEndIndex == -1 {
		return ""
	}
	uuidStartIndex := strings.LastIndex(line[:uuidEndIndex] , ")") + 1
	return line[uuidStartIndex : uuidEndIndex]
}

func changeTaskStateAtLine(line string, toState string) (string){
	toStateValue := "9"
	switch toState {
	case "create": toStateValue = "0"
	case "progress": toStateValue = "1"
	case "suspend": toStateValue = "2"
	case "cancel": toStateValue = "3"
	case "complete": toStateValue = "4"
	}
	
	return setMetadataValue(line, "T", toStateValue)
}

func setMetadataValue(line string, key string, value string) (string) {
	// Not a metadata line, ignore
	if string(line[0]) != "(" {
		return ""
	}
	uuidEndIndex := strings.Index(line, ">");
	// this case should not happen
	if uuidEndIndex == -1 {
		return ""
	}
	uuidStartIndex := strings.LastIndex(line[:uuidEndIndex] , ")") + 1
	metadataStartIndex := strings.Index(line[:uuidStartIndex], "(T-")
	return line[:metadataStartIndex + 3] + value + line[metadataStartIndex + 4:]
}