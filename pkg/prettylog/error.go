package prettylog

import (
	"fmt"
	"log"
	"os"

	"github.com/rhodeon/sniphub/pkg/colors"
)

// Logger which prints message to standard error.
// Also displays the date, time and relative path of the affected file.
var errorLogger = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

// Prints error log and resets text colour afterwards.
func colorizeError(logError func(), error ...interface{}) {
	fmt.Fprint(os.Stderr, colors.Red)
	logError()
	fmt.Fprint(os.Stderr, colors.Reset)
}

// Equivalent to fmt.Print.
func Error(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Print(error...)
	},
	)
}

// Equivalent to fmt.PrintF.
func ErrorF(format string, error ...interface{}) {
	colorizeError(func() {
		errorLogger.Printf(format, error...)
	})
}

// Equivalent to fmt.PrintLn.
func ErrorLn(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Println(error...)
	})
}

// Logs the error then exits the program.
func FatalError(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Fatal(error...)
	})
}
