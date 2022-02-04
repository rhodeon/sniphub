package logging

import (
	"fmt"
	"log"
	"os"

	colors "github.com/rhodeon/sniphub/pkg/color"
)

var infoLog = log.New(os.Stdout, colors.Yellow+"INFO:\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, colors.Red+"ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

// Prints info log and resets text colour afterwards
func LogInfo(message ...interface{}) {
	infoLog.Println(message...)
	fmt.Print(colors.Reset)
}

// Prints error log and resets text colour afterwards
func LogError(error ...interface{}) {
	errorLog.Fatal(error...)
	fmt.Print(colors.Reset)
}
