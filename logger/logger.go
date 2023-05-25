package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Logger is the global logger for the blockchain alert system.
var Logger = InitLogger("")

// InitLogger initializes the global Logger with logs written to the logFile.
// If logFile is empty, logs will be written to os.Stdout.
func InitLogger(logFile string) *log.Logger {
	var err error

	writer := os.Stdout
	if logFile != "" {
		if _, err := os.Stat(logFile); os.IsNotExist(err) {
			tmpStrings := strings.Split(logFile, "/")
			if len(tmpStrings) > 1 {
				directory := strings.Replace(logFile, tmpStrings[len(tmpStrings)-1], "", -1)
				err = os.MkdirAll(directory, os.ModePerm)
				if err != nil {
					fmt.Printf("make directory %v error: %v\n", directory, err)
					os.Exit(1)
				}
			}
		}

		writer, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Error opening file:", err)
			os.Exit(1)
		}
	}

	tmpLogger := log.New(writer, "", log.Ldate|log.Ltime|log.Lshortfile)

	return tmpLogger
}
