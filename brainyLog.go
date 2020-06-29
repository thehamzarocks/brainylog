package main

import (
	"fmt"
	"os"
	"strings"
)

const defaultFilePath = "log.bl"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Invalid usage. Use one of the v, a, g commands")
		return
	}
	command := args[0]
	switch command {
	case "v":
		showBrainyLogVersion()
	case "a":
		if len(args) < 2 {
			fmt.Println("Add epects at least one parameter!")
			return
		}
		processBrainyLogWrite(args)
	case "g":
		processGetCommand(args[1:])
	case "t":
		processTask(args)
	}
}

func isFlag(command string, token string) bool {
	switch command {
	case "g":
		return token == "nm"
	}
	panic("Invalid command!")
}

func isSingleValuedKey(command string, token string) bool {
	switch command {
	case "g":
		return token == "t"
	}
	panic("Invalid command!")
}

func isMultiValuedKey(command string, token string) bool {
	switch command {
	case "g":
		return token == "l"
	}
	panic("Invalid command!")
}

func isValidValueForKey(command string, key string, value string) bool {
	if key == "t" {
		return key != "create" && key != "progress" && key != "suspend" && key != "cancel" && key != "complete"
	}
	return true
}

func processGetCommand(args []string) {
	commandMap := make(map[string]string)
	remainingArgs := args
	tokenTypeToLookFor := "key/flag"
	currentKeyOrFlag := ""
	hasRemaining := true
	for hasRemaining {
		if tokenTypeToLookFor == "key/flag" {
			currentKeyOrFlag = remainingArgs[0]
			if len(remainingArgs) == 1 {
				hasRemaining = false
			} else {
				remainingArgs = remainingArgs[1:]
			}
			if isFlag("g", currentKeyOrFlag) {
				commandMap[currentKeyOrFlag] = ""
				tokenTypeToLookFor = "key/flag"
				continue
			}
			if isSingleValuedKey("g", currentKeyOrFlag) {
				tokenTypeToLookFor = "singleValue"
				continue
			}
			if isMultiValuedKey("g", currentKeyOrFlag) {
				tokenTypeToLookFor = "multiValue"
				continue
			}
			fmt.Println("Invalid key/flag " + currentKeyOrFlag + "!")
			return
		}

		if tokenTypeToLookFor == "singleValue" {
			currentValue := remainingArgs[0]
			if len(remainingArgs) == 1 {
				hasRemaining = false
			} else {
				remainingArgs = remainingArgs[1:]
			}
			if !isValidValueForKey("g", currentKeyOrFlag, currentValue) {
				fmt.Println("Invalid value " + currentValue + " for key " + currentKeyOrFlag + "!")
				return
			}
			commandMap[currentKeyOrFlag] = currentValue
			tokenTypeToLookFor = "key/flag"
		}

		if tokenTypeToLookFor == "multiValue" {
			tokenTypeToLookFor = "none"
			commandMap[currentKeyOrFlag] = strings.Join(remainingArgs, " ")
			break
		}
	}

	if tokenTypeToLookFor == "singleValue" || tokenTypeToLookFor == "multiValue" {
		fmt.Println("Expected value for key " + currentKeyOrFlag + "!")
		return
	}

	processBrainyLogRead(commandMap)
}
