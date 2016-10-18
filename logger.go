package mlog

import (
	//	"fmt"
	"io"
	"log"
	"os"
)

const (
	// DefFlags define the default flags
	DefFlags = log.LstdFlags

	// DefPrefix defines the log message prefix
	DefPrefix = ""

	// DefFormat defines default message format
	DefFormat = "%-10s %s"

	// DefPriority defines the default priority (severity) level
	DefPriority = LogInfo
)

// Logger is a more flexible version of the standard lib logger.
type Logger struct {

	// a single instance of the standard lib logger
	log.Logger

	// log handler defining severity and format
	Handler

	// channel to send messages to this logger
	Msgch chan string

	// a channel used to stop the logger
	Stopch chan bool
}

// CreateLogger creates a new Logger instance with all custom info.
func CreateLogger(output io.Writer, prio Priority, prefix, format string, flags int) *Logger {

	l := log.New(output, prefix+" ", flags)
	logger := &Logger{*l, *NewHandler(format, prio), make(chan string), make(chan bool)}

	// run the logging goroutine for this logger
	go func(l *Logger) {

		var msg string

		for {
			select {

			// we receive the new log message...
			case msg = <-l.Msgch:
				l.Print(msg)

			// when the stop signal is received, both channels are closed and goroutine gracefully dies...
			case <-l.Stopch:
				close(l.Msgch)
				close(l.Stopch)
				return
			}
		}
	}(logger)

	return logger
}

// CreateFileLogger creates a new Logger instance that writes messages into file, with all custom info.
func CreateFileLogger(path string, prio Priority, prefix, format string, flags int) (*Logger, error) {

	// create new file logger with 'append' flag
	fin, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	return CreateLogger(fin, prio, prefix, format, flags), err
}

// ClearFileLogger clears the logfile with the fiven path.
// The function actually just deletes the existing file and creates new logger with the same attributes (including filename).
func ClearFileLogger(path string, logger *Logger) (*Logger, error) {

	if err := os.Remove(path); err != nil {
		return nil, err
	}
	return CreateFileLogger(path, logger.Priority, logger.Prefix(), logger.Format, logger.Flags())
}

// NewLogger creates a new Logger instance with all default info.
// The output must satisfy the io.Writer interface.
func NewLogger(output io.Writer) *Logger {
	return CreateLogger(output, DefPriority, DefPrefix, DefFormat, DefFlags)
}

// SetFormat defines the custom log message format.
func (l *Logger) SetFormat(fmt string) { l.Format = fmt }

// SetPriority defines the severity for the Logger.
func (l *Logger) SetPriority(prio Priority) { l.Priority = prio }
