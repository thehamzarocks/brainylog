package main

import (
	"fmt"
	"os"
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
	if isTask {
		addTaskLog(log)
		return
	}
	addInfoLog(log)
	return
}

func addInfoLog(log string) {
	filename := defaultFilePath
	line := processLine("info", log)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(line + "\n"); err != nil {
		panic(err)
	}
	fmt.Println("Info logged:\n\n" + line + "\n")
}

func addTaskLog(log string) {
	filename := defaultFilePath
	line := processLine("task", log)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(line + "\n"); err != nil {
		panic(err)
	}
	fmt.Println("Task Logged:\n\n" + line + "\n")
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
