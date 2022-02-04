package prettylog

import (
	"fmt"
	"log"
	"os"

	"github.com/rhodeon/sniphub/pkg/colors"
)

// Logger which prints message to standard output.
// Also displays the date and time.
var infoLogger = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

// Prints info log and resets text colour afterwards.
func colorizeInfo(logInfo func(), info ...interface{}) {
	fmt.Fprint(os.Stdout, colors.Yellow)
	logInfo()
	fmt.Fprint(os.Stdout, colors.Reset)
}

// Equivalent to fmt.Print.
func Info(info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Print(info...)
	}, info...)
}

// Equivalent to fmt.Printf.
func InfoF(format string, info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Printf(format, info...)
	})
}

// Equivalent to fmt.Println.
func InfoLn(format string, info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Println(info...)
	})
}
