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
			if currentKeyOrFlag == "nm" {
				commandMap["nm"] = ""
				tokenTypeToLookFor = "key/flag"
				continue
			}
			if currentKeyOrFlag == "t" {
				tokenTypeToLookFor = "singleValue"
				continue
			}
			if currentKeyOrFlag == "l" {
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
			if currentKeyOrFlag == "t" {
				if currentValue != "create" && currentValue != "progress" && currentValue != "suspend" && currentValue != "cancel" && currentValue != "complete" {
					fmt.Println("Invalid value " + currentValue + " for key t!")
					return
				}
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
