package logger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInitLogger_EmptyLogFile(t *testing.T) {
	Logger = InitLogger("")

	Logger.Printf("Hello World!\n")
}

func TestInitLogger_NonEmptyLogFile(t *testing.T) {
	logFile := "example_log.log"

	// delete the logFile before init
	_, err := os.Stat(logFile)
	if !os.IsNotExist(err) {
		err = os.Remove(logFile)
		if err != nil {
			panic(err)
		}
	}

	// init the logger
	Logger = InitLogger(logFile)

	msg := "Hello World!\n"
	Logger.Printf(msg)

	f, err := os.Open(logFile)
	if err != nil {
		panic(fmt.Sprintf("cannot open file %v: %v", logFile, err))
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("cannot read file %v: %v", logFile, err))
	}
	assert.Equal(t, true, strings.Contains(string(data), msg))
}
