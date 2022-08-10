package logger

import (
	"fmt"
	"log"
)

var log_level int = 0

// Set LogLevel for logger:
// [0] Errors are getting logged,
// [1] Errors and Warnings are getting logged,
// [2] Errors, Warnings and Infos are getting logged
func SetLogLevel(level int) {
	if level < 3 && level >= 0 {
		log_level = level
	} else {
		panic(fmt.Sprintf("Unknown LogLevel \"%d\". Possible values between 0-2", level))
	}
}

// Log info message *message* if log_level is set to 2
func Info(message string) {
	if log_level == 2 {
		info := string(colorWhite) + "INFO" + string(colorReset)
		log.Printf("[  %s   ] %v", info, message)
	}
}

// Log warning message *message* if log_level is at least set to 1
func Warning(message string) {
	if log_level >= 1 {
		warning := string(colorYellow) + "WARNING" + string(colorReset)
		log.Printf("[ %s ] %v", warning, message)
	}
}

// Log error message *message*
func Error(message string) {
	if log_level >= 0 {
		err := string(colorRed) + "ERROR" + string(colorReset)
		log.Printf("[  %s  ] %v", err, message)
	}
}
