package utils

import (
	"fmt"
	"log"
	"os"
)

type ILogger interface {
	Success(msg string)
	Successf(format string, v ...interface{})
	Info(msg string)
	Infof(format string, v ...interface{})
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Error(msg string)
	Errorf(format string, v ...interface{})
}

// Logger defines a simple structured logger.
type Logger struct {
	SuccessLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarnLogger    *log.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{
		SuccessLogger: log.New(os.Stdout, "\033[32mSUCCESS\033[0m: ", log.LstdFlags|log.Lshortfile),
		InfoLogger:    log.New(os.Stdout, "\033[34mINFO\033[0m: ", log.LstdFlags|log.Lshortfile),
		ErrorLogger:   log.New(os.Stderr, "\033[31mERROR\033[0m: ", log.LstdFlags|log.Lshortfile),
		WarnLogger:    log.New(os.Stderr, "\033[33mWARN\033[0m: ", log.LstdFlags|log.Lshortfile),
	}
}

// Success logs success messages.
func (l *Logger) Success(msg string) {
	l.SuccessLogger.Output(2, fmt.Sprintln(msg))
}

// Successf logs formatted success messages.
func (l *Logger) Successf(format string, v ...interface{}) {
	l.SuccessLogger.Output(2, fmt.Sprintf(format, v...))
}

// Info logs informational messages.
func (l *Logger) Info(msg string) {
	l.InfoLogger.Output(2, fmt.Sprintln(msg))
}

// Infof logs formatted informational messages.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

// Warn logs warning messages.
func (l *Logger) Warn(msg string) {
	l.WarnLogger.Output(2, fmt.Sprintln(msg))
}

// Warnf logs formatted warning messages.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WarnLogger.Output(2, fmt.Sprintf(format, v...))
}

// Error logs error messages.
func (l *Logger) Error(msg string) {
	l.ErrorLogger.Output(2, fmt.Sprintln(msg))
}

// Errorf logs formatted error messages.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.ErrorLogger.Output(2, fmt.Sprintf(format, v...))
}
