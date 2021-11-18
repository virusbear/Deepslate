package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	name string
	log *log.Logger
}

func NewLogger(name string) *Logger {
	logger := &Logger{name, log.New()}
	initiliazeLogger(logger)
	return logger
}

func (logger Logger) Panic(args ...interface{}) {
	logger.newEntry().Panic(args...)
}

func (logger Logger) Fatal(args ...interface{}) {
	logger.newEntry().Fatal(args...)
}

func (logger Logger) Error(args ...interface{}) {
	logger.newEntry().Error(args...)
}

func (logger Logger) Warn(args ...interface{}) {
	logger.newEntry().Warn(args...)
}

func (logger Logger) Info(args ...interface{}) {
	logger.newEntry().Info(args...)
}

func (logger Logger) Debug(args ...interface{}) {
	logger.newEntry().Debug(args...)
}

func (logger Logger) Trace(args ...interface{}) {
	logger.newEntry().Trace(args...)
}

func (logger Logger) newEntry() *log.Entry {
	return logger.log.WithField("logger", logger.name)
}

func initiliazeLogger(logger *Logger) {
	logger.log.SetOutput(os.Stdout)
	logger.log.SetLevel(log.DebugLevel)
}
