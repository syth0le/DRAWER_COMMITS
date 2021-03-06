package main

import (
	"os"
	"strings"
	"time"
)

func getCurrentDateTime() string {
	dateTime := time.Now().Format("2006-01-02 15:04:05.000000000")
	return "[" + dateTime + "] - "
}

func writeToLogFile(message ...string) bool {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	dateTime := getCurrentDateTime()
	valueToWrite := dateTime + "git " + strings.Join(message, " ") + "\n"

	_, err = file.WriteString(valueToWrite)
	if err != nil {
		panic(err)
	}
	file.Close()
	return true
}

func cleanLogFile() bool {
	if !isCleanLogFile {
		return isCleanLogFile
	}
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString("")
	if err != nil {
		panic(err)
	}
	return isCleanLogFile
}
