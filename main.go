package logger

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

var log_level int = 0
var log_filename string = ""
var ch = make(chan string)

// Set LogLevel for logger:
// [0] Errors are getting logged,
// [1] Errors and Warnings are getting logged,
// [2] Errors, Warnings and Infos are getting logged
func SetLogLevel(level int) {
	if level < 3 && level >= 0 {
		log_level = level
	} else {
		panic(fmt.Sprintf("Unknown Log Level \"%d\". Possible values between 0-2", level))
	}
}

// Set Log Filename for logger and start go routine for listening on log channel
func SetLogFilename(name string) {
	re, _ := regexp.Compile(` (\.{0,2}\/)?(\.{1,2}\/)*(\w+\/)*[\w-() _]+[\w-() _\.]* `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", name))) {
		panic("Incorrect filename for output file")
	}
	log_filename = name
	go writeLogsToFile()
}

// Receive logs from log channel and write them to log_file
func writeLogsToFile() {
	log_file, err := os.OpenFile(log_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Unable to create/open file %s for writing", log_filename))
	}

	for log := range ch {
		log_file.WriteString(log + "\n")
	}

	defer log_file.Close()
}

// Log info message *message* if log_level is set to 2
func Info(message string) {
	if log_level == 2 {
		header := time.Now().Format("2006/02/01 03:04:05")
		if log_filename != "" {
			ch <- fmt.Sprintf("%s [  INFO   ] %v", header, message)
		} else {
			info := string(colorWhite) + "INFO" + string(colorReset)
			fmt.Printf("%s [  %s   ] %v\n", header, info, message)
		}
	}
}

// Log warning message *message* if log_level is at least set to 1
func Warning(message string) {
	if log_level >= 1 {
		header := time.Now().Format("2006/02/01 03:04:05")
		if log_filename != "" {
			ch <- fmt.Sprintf("%s [ WARNING ] %v", header, message)
		} else {
			warning := string(colorYellow) + "WARNING" + string(colorReset)
			fmt.Printf("%s [  %s   ] %v\n", header, warning, message)
		}
	}
}

// Log error message *message*
func Error(message string) {
	if log_level >= 0 {
		header := time.Now().Format("2006/02/01 03:04:05")

		if log_filename != "" {
			ch <- fmt.Sprintf("%s [  ERROR  ] %v", header, message)
		}

		err := string(colorRed) + "ERROR" + string(colorReset)
		fmt.Printf("%s [  %s  ] %v\n", header, err, message)
	}
}

// Log fatal message *message* and exit program afterwards
func Fatal(message string) {
	if log_level >= 0 {
		header := time.Now().Format("2006/02/01 03:04:05")

		if log_filename != "" {
			ch <- fmt.Sprintf("%s [  FATAL  ] %v", header, message)
		}

		err := string(colorBackgroundRed) + "FATAL" + string(colorReset)
		fmt.Printf("%s [  %s  ] %v\n", header, err, message)
	}
	os.Exit(1)
}
