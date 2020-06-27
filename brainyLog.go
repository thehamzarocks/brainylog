package main

import (
	"fmt"
	"os"
)

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
		logType, log := processAddArgs(args)
		if logType == "noLog" {
			fmt.Println("Invalid usage. Pass in a message to be logged")
			return
		}
		addBrainyLogLog(logType, log)
	case "g":
		searchText, err := processGetArgs(args)
		if (err != nil) {
			fmt.Println("Invalid usage. Please pass in a search text")
			return
		}
		getBrainyLogMatches(searchText)
	}
}