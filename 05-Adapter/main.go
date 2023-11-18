/*
	Suppose we have legacy code written and their is a logger in it.
	Now in new code we have a new logger. But we don't want to rewrite whole legacy application according to new logger.
	In this case we will create an adapter that will allow us to use legacy logger with new logger.
*/

package main

import "fmt"

type LegacyLogger struct{}

func (l *LegacyLogger) Log(msg string) {
	fmt.Println("Legacy Logger: ", msg)
}

// New Logger uses an interface to log a message
type NewLogger interface {
	Log(msg string)
}

type NewSystemLogger struct{}

func (l *NewSystemLogger) Log(msg string) {
	fmt.Println("New System Logger: ", msg)
}

// Now if we want to use New Logger with Legacy Logger than we can use an adapter for it.
type LoggerAdapter struct {
	legacyLogger *LegacyLogger
}

func (l *LoggerAdapter) Log(msg string) {
	l.legacyLogger.Log(msg)
}

func logMessage(msg string, logger NewLogger) {
	logger.Log(msg)
}

func main() {
	legacyLogger := &LegacyLogger{}
	loggerAdapter := &LoggerAdapter{legacyLogger: legacyLogger}

	logMessage("Hello", loggerAdapter)
}
