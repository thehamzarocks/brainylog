package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"os"

	"github.com/google/uuid"
)

func processAddArgs(args []string) (string, string) {
	logType := "info"
	if len(args) < 2 {
		return "noLog", ""
	}
	log := strings.Join(args[1:], " ")
	return logType, log
}

func addBrainyLogLog(logType string, log string) {
	filename := "bin/log.bl"
	line := processLine(logType, log)
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

func processLine(logtype string, log string) string {
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
	str.WriteString(uuid.New().String())
	str.WriteString(">")
	str.WriteString(log)

	return str.String()
}

