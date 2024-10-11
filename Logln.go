package logln

import (
	"fmt"
	"os"
	"time"
)

// fileName is the name of the log file. It is initialized with the current date and time in the format YYYY_MM_DD-HH_mm_SS.log
var fileName = getDateStringFileFmt() + ".log"

// logFile is a pointer to the log file. It is initialized to nil and is set to the opened log file in the Init function.
var logFile *os.File

// Init initializes the log file by opening it with the name fileName. If the file does not exist, it is created.
// If there is an error opening the file, an error message is printed to the console and the program exits with status code 1.
func Init() {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
	logFile = file
}

// Close closes the log file
func Close() error {
	if logFile != nil {
		return logFile.Close()
	}
	return nil
}

// getDateStringLogFmt returns the current date and time in the format YYYY/MM/DD HH:mm:SS
func getDateStringLogFmt() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// getDateStringFileFmt returns the current date and time in the format YYYY_MM_DD-HH_mm_SS
func getDateStringFileFmt() string {
	return time.Now().Format("2006_01_02-15_04_05")
}

// levelStr returns a string representation of the log level based on the provided integer.
// It takes an integer 'level' as input and returns a string representing the log level.
// The function uses a switch statement to map the integer values to their corresponding string representations.
// If the input integer does not match any of the cases, an empty string is returned.
func levelStr(level int) string {
	switch level {
	case 0:
		return "INFO" // If the level is 0, return "INFO"
	case 1:
		return "WARNING" // If the level is 1, return "WARNING"
	case 2:
		return "ERROR" // If the level is 2, return "ERROR"
	case 3:
		return "FATAL" // If the level is 3, return "FATAL"
	case 5:
		return "DEBUG" // If the level is 5, return "DEBUG"
	default:
		return "" // If the level does not match any case, return an empty string
	}
}

// PrintWarningOrSuccessIfNotOk logs an error message if the 'ok' parameter is false,
// otherwise logs a success message.
func PrintWarningOrSuccessIfNotOk(ok bool, msg string, successLevel int, isSuccessSilent bool) {
	if !ok {
		warnIfNotOk(ok, msg)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintErrorOrSuccessIfNotOk logs an error message if the 'ok' parameter is false,
// otherwise logs a success message.
func PrintErrorOrSuccessIfNotOk(ok bool, msg string, successLevel int, isSuccessSilent bool) {
	if !ok {
		errorIfNotFalse(ok, msg)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintFatalOrSuccessIfNotOk logs a fatal error message if the 'ok' parameter is false,
// otherwise logs a success message. If a fatal error is logged, the program will exit with status code 1.
func PrintFatalOrSuccessIfNotOk(ok bool, msg string, successLevel int, isSuccessSilent bool) {
	if !ok {
		fatalIfNotFalse(ok, msg)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintPanicOrSuccessIfNotOk logs a panic message if the 'ok' parameter is false,
// otherwise logs a success message. If a panic is logged, the program will crash.
func PrintPanicOrSuccessIfNotOk(ok bool, msg string, successLevel int, isSuccessSilent bool) {
	if !ok {
		panicIfNotFalse(ok, msg)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintWarningOrSuccess logs a warning message if the 'err' parameter is not nil,
// otherwise logs a success message.
func PrintWarningOrSuccess(msg string, err error, successLevel int, isSuccessSilent bool) {
	if err != nil {
		warningWithError(msg, err)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintErrorOrSuccess logs an error message if the 'err' parameter is not nil,
// otherwise logs a success message.
func PrintErrorOrSuccess(msg string, err error, successLevel int, isSuccessSilent bool) {
	if err != nil {
		printError(msg, err)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintPanicOrSuccess logs a panic message if the 'err' parameter is not nil,
// otherwise logs a success message. If a panic is logged, the program will crash.
func PrintPanicOrSuccess(msg string, err error, successLevel int, isSuccessSilent bool) {
	if err != nil {
		panicError(msg, err)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// PrintFatalOrSuccess logs a fatal error message if the 'err' parameter is not nil,
// otherwise logs a success message. If a fatal error is logged, the program will exit with status code 1.
func PrintFatalOrSuccess(msg string, err error, successLevel int, isSuccessSilent bool) {
	if err != nil {
		fatalError(msg, err)
	} else {
		printSuccess(msg, successLevel, isSuccessSilent)
	}
}

// errorIfNotOk logs an error message if the 'ok' parameter is false.
func errorIfNotOk(ok bool, msg string) {
	errorIfNotFalse(ok, msg)
}

// warnIfNotOk logs a warning message if the 'ok' parameter is false.
func warnIfNotOk(ok bool, msg string) {
	warnIfNotFalse(ok, msg)
}

// fatalIfNotOk logs a fatal error message if the 'ok' parameter is false.
// The program will exit with status code 1.
func fatalIfNotOk(ok bool, msg string) {
	fatalIfNotFalse(ok, msg)
}

// panicIfNotOk logs a panic message if the 'ok' parameter is false.
// The program will then panic
func panicIfNotOk(ok bool, msg string) {
	panicIfNotFalse(ok, msg)
}

// warningWithError logs a warning message with the given message and error.
func warningWithError(msg string, err error) {
	Logln(fmt.Sprintf("%s: %s", msg, err.Error()), 1, false)
}

// warning logs a simple warning message with the given message.
func warning(msg string, err error) {
	Logln(fmt.Sprintf("%s: %s", msg, err.Error()), 1, false)
}

// printError logs an error message with the given message and error.
func printError(msg string, err error) {
	Logln(fmt.Sprintf("Error occurred %s: %s", msg, err.Error()), 2, false)
}

// fatalError logs a fatal error message with the given message and error,
// then exits the program with status code 1.
func fatalError(msg string, err error) {
	Logln(fmt.Sprintf("Fatal error encountered %s: %s", msg, err.Error()), 3, false)
	os.Exit(1)
}

// panicError logs a panic message with the given message and error.
// The program will crash.
func panicError(msg string, err error) {
	Logln(fmt.Sprintf("Panic %s: %s", msg, err.Error()), 3, false)
}

// errorIfNotFalse logs an error message if the 'b' parameter is false.
func errorIfNotFalse(b bool, msg string) {
	switch b {
	case false:
		Logln(fmt.Sprintf("Error occurred %s: %t", msg, b), 2, false)
	}
}

// warnIfNotFalse logs a warning message if the 'b' parameter is false.
func warnIfNotFalse(b bool, msg string) {
	switch b {
	case false:
		Logln(fmt.Sprintf("%s: %t", msg, b), 1, false)
	}
}

// fatalIfNotFalse logs a fatal error message if the 'b' parameter is false.
// The program will exit with status code 1.
func fatalIfNotFalse(b bool, msg string) {
	switch b {
	case false:
		Logln(fmt.Sprintf("Fatal error occurred %s: %t", msg, b), 3, false)
	}
}

// panicIfNotFalse logs a panic message if the 'b' parameter is false.
// The program will crash.
func panicIfNotFalse(b bool, msg string) {
	switch b {
	case false:
		Logln(fmt.Sprintf("Panic %s: %t", msg, b), 3, false)
	}
}

// printSuccess logs a success message with the given message, level, and silent flag.
func printSuccess(msg string, successLevel int, isSilent bool) {
	Logln(fmt.Sprintf("Success %s.", msg), successLevel, isSilent)
}

// ManualLogf writes the given text to the console and appends it to the log file with the current date and time. It doesn't append a newline to the provided string.
func ManualLogf(text string, level int, isSilent bool) error {
	levelStr := levelStr(level)

	date := getDateStringLogFmt()
	msg := ""
	if levelStr != "" {
		msg = date + " " + levelStr + " " + text
	} else {
		msg = date + " " + text
	}

	if !isSilent {
		fmt.Print(msg)
	}
	if logFile != nil {
		_, err := logFile.WriteString(msg)
		return err
	}
	return nil
}

// Printf writes the given text to the console and appends it to the log file. It doesn't append a newline to the provided string.
func Printf(text string, level int, isSilent bool) error {
	if !isSilent {
		fmt.Print(text)
	}
	if logFile != nil {
		_, err := logFile.WriteString(text)
		return err
	}
	return nil
}

// Logln writes the given line to the console and appends it to the log file with the current date and time. It appends a newline to the provided string.
func Logln(line string, level int, isSilent bool) error {
	// Get the level string based on the level integer
	levelStr := levelStr(level)
	// Get the current date and time in the format YYYY/MM/DD HH:mm:SS
	date := getDateStringLogFmt()
	// Initialize an empty string for the message
	msg := ""

	// If the level string is not empty, include it in the message
	if levelStr != "" {
		msg = date + " " + levelStr + " " + line
	} else {
		// If the level string is empty, do not include it in the message
		msg = date + " " + line
	}

	// If the log is not silent
	if !isSilent {
		// Print the line to the console
		fmt.Println(msg + "\n")
	}

	// If the log file is not nil
	if logFile != nil {
		// Write the message to the log file, appending a newline
		_, err := logFile.WriteString(msg + "\n")
		return err
	}
	// If the log file is nil, return nil
	return nil
}
