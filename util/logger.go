package util

import (
	"log"
	"os"
)

var Logger *CustomLogger

type CustomLogger struct {
	logger *log.Logger
}

func NewCustomLogger(logFilePath string) (*CustomLogger, error) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}, nil
}

func (c *CustomLogger) LogInfo(message string) {
	c.logger.Println("[INFO]", message)
}

func (c *CustomLogger) LogWarning(message string) {
	c.logger.Println("[WARNING]", message)
}

func (c *CustomLogger) LogError(message string) {
	c.logger.Println("[ERROR]", message)
}
