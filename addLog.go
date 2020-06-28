package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"os"

	"github.com/google/uuid"
)

func processBrainyLogWrite(args []string) {
	logType, log := processAddArgs(args)
		if logType == "noLog" {
			fmt.Println("Invalid usage. Pass in a message to be logged")
			return
		}
		if logType == "info" {
			addInfoLog(logType, log)
		}
		if logType == "task" {
			addTaskLog(logType, log)
		}
	}

func processAddArgs(args []string) (string, string) {
	if len(args) < 2 {
		return "noLog", ""
	}

	if args[1] == "-t" {
		log := strings.Join(args[2:], " ")
		return "task", log
	}

	log := strings.Join(args[1:], " ")
	return "info", log
}

func addInfoLog(logType string, log string) {
	filename := "bin/log.bl"
	line := processLine(logType, log, "info")
	fmt.Println(filename, line)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(line + "\n"); err != nil {
		panic(err)
	}
}

func addTaskLog(logType string, log string) {
	filename := "bin/log.bl"
	line := processLine(logType, log, "task")
	fmt.Println(filename, line)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(line + "\n"); err != nil {
		panic(err)
	}
}

func processLine(logtype string, log string, logType string) string {
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
	
	if logType == "task" {
		str.WriteString("(T-0)")
	}

	str.WriteString(uuid.New().String())
	str.WriteString(">")
	str.WriteString(log)

	return str.String()
}

