package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	config.Encoding = "json"

	newLogger, err := config.Build()
	if err != nil {
		fmt.Printf("failed to create logger: %v", err)
		os.Exit(1)
	}

	logger = newLogger.Sugar()
}

func Error(err error, keyArgsValues ...any) {
	logger.Errorw(fmt.Sprint(err), keyArgsValues...)
}

func Info(message string, keyArgsValues ...any) {
	logger.Infow(message, keyArgsValues...)
}

func Debug(message string, keyArgsValues ...any) {
	logger.Debugw(message, keyArgsValues...)
}

func Fatal(message string, keyArgsValues ...any) {
	logger.Fatalw(message, keyArgsValues...)
}
