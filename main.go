package logger

import "log"

func LogInfo(m string) {
	log.Printf("INFO - %v", m)
}

func LogWarning(m string) {
	log.Printf("WARNING - %v", m)
}

func LogError(m string) {
	log.Printf("ERROR - %v", m)
}
