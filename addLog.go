package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func processBrainyLogWrite(commandMap map[string]string) {
	log, containsLog := commandMap["l"]
	if !containsLog {
		fmt.Println("Please pass in a message to be logged!")
		return
	}
	if strings.Index(log, "\n") != -1 {
		fmt.Println("Only single-line logs are permitted!")
		return
	}

	_, isTask := commandMap["t"]

	temporalPosition := commandMap["n"]

	argsMap := make(map[string]string)
	argsMap["log"] = log
	argsMap["isTask"] = strconv.FormatBool(isTask)
	argsMap["temporalPosition"] = temporalPosition
	processFile(addLog, argsMap)
	return
}

func addLog(lines []string, argsMap map[string]string) (linesToWrite []string, shouldWriteLines bool) {
	log := argsMap["log"]
	// temporalPosition := argsMap["temporalPosition"]

	isTask, err := strconv.ParseBool(argsMap["isTask"])
	if err != nil {
		fmt.Println("Unexpected error while parsing task flag!")
		return lines, false
	}

	var line string
	if isTask {
		line = processLine("task", log)
	} else {
		line = processLine("info", log)
	}

	lines = append(lines, line)

	if isTask {
		fmt.Println("Task logged:\n\n" + line + "\n")
	} else {
		fmt.Println("Info logged:\n\n" + line + "\n")
	}

	return lines, true
}

func processLine(logType string, log string) string {
	t := time.Now()
	timestamp := t.Format(time.RFC3339)
	epoch := t.UnixNano() / 1e6

	var str strings.Builder
	str.WriteString("(")
	str.WriteString(timestamp)
	str.WriteString(")")
	str.WriteString("[")
	str.WriteString(strconv.FormatInt(epoch, 10))
	str.WriteString(".0")
	str.WriteString("]")

	str.WriteString("(S-00)")

	if logType == "task" {
		str.WriteString("(T-0)")
	}

	str.WriteString(uuid.New().String())
	str.WriteString(">")
	str.WriteString(log)

	return str.String()
}
