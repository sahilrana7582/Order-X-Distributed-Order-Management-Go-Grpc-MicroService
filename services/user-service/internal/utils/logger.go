package utils

import (
	"log"
	"os"
)

func InitLogger() (infoLog *log.Logger, errorLog *log.Logger) {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
