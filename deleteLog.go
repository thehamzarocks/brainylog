package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func processDeleteLog(commandMap map[string]string) {
	matches := 0
	uuidMatcher, containsUUID := commandMap["u"]
	positionMatcher, containsPosition := commandMap["n"]
	if containsUUID {
		matches++
	}
	if containsPosition {
		matches++
	}
	if matches != 1 {
		fmt.Println("Delete needs exactly one matcher!")
		return
	}

	if containsUUID {
		deleteLine(uuidMatcher)
		return
	}
	if containsPosition {
		lineUUID, err := getUUIDFromTemporaryPositionalNumber(positionMatcher)
		if err != nil {
			fmt.Println(err.Error())
		}
		deleteLine(lineUUID)
		return
	}
}

func deleteLine(uuid string) {
	filename := defaultFilePath
	// fmt.Println(filename, line)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	if len(lines) == 0 {
		fmt.Println("Empty file!")
		return
	}
	lines = lines[:len(lines)-1]

	for i, line := range lines {
		if getUUID(line) == uuid {
			changedLine := setMetadataValue(line, "S", "01")
			lines[i] = changedLine
			fmt.Println("Line deleted:")
			fmt.Println(changedLine)
			break
		}
	}

	output := strings.Join(lines, "\n")
	output += "\n"
	err = ioutil.WriteFile(defaultFilePath, []byte(output), 0644)

	if err != nil {
		fmt.Println("Something went wrong while deleting line")
	}
}
