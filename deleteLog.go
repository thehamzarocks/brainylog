package main

import (
	"fmt"
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

	argsMap := make(map[string]string)

	if containsUUID {
		argsMap["uuid"] = uuidMatcher
		processFile(deleteLine, argsMap)
		return
	}
	if containsPosition {
		lineUUID, err := getUUIDFromTemporaryPositionalNumber(positionMatcher)
		if err != nil {
			fmt.Println(err.Error())
		}
		argsMap["uuid"] = lineUUID
		processFile(deleteLine, argsMap)
		return
	}
}

func deleteLine(lines []string, argsMap map[string]string) (writeBack bool) {

	uuid := argsMap["uuid"]

	for i, line := range lines {
		if getUUID(line) == uuid {
			changedLine := setMetadataValue(line, "S", "01")
			lines[i] = changedLine
			fmt.Println("Line deleted:")
			fmt.Println("")
			fmt.Println(changedLine)
			fmt.Println("")
			break
		}
	}

	return true
}
