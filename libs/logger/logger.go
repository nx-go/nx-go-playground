package logger

import "log"

func Info(message string) {
	log.Println("[INFO] " + message)
}

func Warn(message string) {
	log.Println("[WARN] " + message)
}

func Debug(message string) {
	log.Println("[DEBUG] " + message)
}

func Error(message string) {
	log.Println("[ERROR] " + message)
}
