package main

import (
	"fmt"
	"os"
	"strings"
)

func processAddArgs(args []string) (string, string) {
	logType := "info"
	if len(args) < 2 {
		return "noLog", ""
	}
	log := strings.Join(args[1:], " ")
	return logType, log
}

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
	}
}

func showBrainyLogVersion() {
	fmt.Println("BrainyLog version 0.1")
	fmt.Println("Thanks for being an early adopter! Unless you're using an old version on purpose.")
}

func addBrainyLogLog(logType string, log string) {
	fmt.Println(logType)
	fmt.Println(log)
}
