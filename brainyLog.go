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
		processCommand("a", args[1:])
	case "g":
		if len(args) < 2 {
			fmt.Println("Get epects at least one parameter!")
			return
		}
		processCommand("g", args[1:])
	case "t":
		if len(args) < 2 {
			fmt.Println("Get epects at least one parameter!")
			return
		}
		processCommand("t", args[1:])
	}
}

func isFlag(command string, token string) bool {
	switch command {
	case "a":
		return token == "t"
	case "g":
		return token == "nm"
	case "t":
		return false
	}
	panic("Invalid command!")
}

func isSingleValuedKey(command string, token string) bool {
	switch command {
	case "a":
		return false
	case "g":
		return token == "t"
	case "t":
		return token == "t" || token == "n" || token == "u"
	}
	panic("Invalid command!")
}

func isMultiValuedKey(command string, token string) bool {
	switch command {
	case "a":
		return token == "l"
	case "g":
		return token == "l"
	case "t":
		return false
	}
	panic("Invalid command " + command + "!")
}

func isValidValueForKey(command string, key string, value string) bool {
	switch command {
	case "a":
		return true
	case "g":
		if key == "t" {
			return !(value != "create" && value != "progress" && value != "suspend" && value != "cancel" && value != "complete")
		}
		return true
	case "t":
		if key == "t" {
			return !(value != "create" && value != "progress" && value != "suspend" && value != "cancel" && value != "complete")
		}
		return true
	}
	panic("Invalid command " + command + "!")
}

func processCommand(command string, args []string) {
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
			if isFlag(command, currentKeyOrFlag) {
				commandMap[currentKeyOrFlag] = ""
				tokenTypeToLookFor = "key/flag"
				continue
			}
			if isSingleValuedKey(command, currentKeyOrFlag) {
				tokenTypeToLookFor = "singleValue"
				continue
			}
			if isMultiValuedKey(command, currentKeyOrFlag) {
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
			if !isValidValueForKey(command, currentKeyOrFlag, currentValue) {
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

	executeCommand(command, commandMap)
}

func executeCommand(command string, commandMap map[string]string) {
	switch command {
	case "a":
		processBrainyLogWrite(commandMap)
	case "g":
		processBrainyLogRead(commandMap)
	case "t":
		processTask(commandMap)
	default:
		panic("Unknown command " + command + "!")
	}
}
