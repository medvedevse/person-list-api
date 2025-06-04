package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("Can't initialize zap logger: %v", err)
		return nil
	}
	return logger
}
