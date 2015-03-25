// Log is the custom package that is used instead of go's standard 'log' package.
// It provides another layer of abstraction for logging to ease the pain of logging to more destinations simultaneouly.
package log

import (
	"fmt"
	"log"
	"strings"
	//    "log/syslog"
)

// Log is a collection (a slice) of Loggers.
// This advanced Log type is used to log messages simultaneously on different loggers. Like STDIN and logfile at the same time or
// even several logfiles (with different message format and severity) at the same time.
type Log struct {
	loggers []*Logger
}

// NewLog creates an empty Log instance.
func NewLog() *Log { return &Log{make([]*Logger, 0)} }

// AddLogger appends a new Logger to the Log.
func (l *Log) AddLogger(logger *Logger) { l.loggers = append(l.loggers, logger) }

// The logmsg method is a private helper method for logging messages with different severities.
func (l *Log) logmsg(prio Priority, msg string) {

	for _, lg := range l.loggers {

		// send message to channel only when severity is smaller or equal to defined...
		if prio <= lg.Priority {
			lg.Msgch <- fmt.Sprintf(lg.Format, strings.ToUpper(prio.String()), msg)
		}
	}
}

// The logmsgf method is a private helper method for logging custom formatted messages with different severities.
func (l *Log) logmsgf(frmt string, prio Priority, v ...interface{}) {

	for _, lg := range l.loggers {

		// send message to channel only when severity is smaller or equal to defined...
		if prio <= lg.Priority {
			//lg.Printf(frmt, v...)
			lg.Msgch <- fmt.Sprintf(frmt, v...)
		}
	}
}

// Debug logs a message with the DEBUG severity.
func (l *Log) Debug(msg string) { l.logmsg(LogDebug, msg) }

// Info logs a message with the INFO severity.
func (l *Log) Info(msg string) { l.logmsg(LogInfo, msg) }

// Notice logs a message with the NOTIC severity.
func (l *Log) Notice(msg string) { l.logmsg(LogNotice, msg) }

// Warning logs a message with the WARN severity.
func (l *Log) Warning(msg string) { l.logmsg(LogWarning, msg) }

// Err logs a message with the ERROR severity.
func (l *Log) Err(msg string) { l.logmsg(LogError, msg) }

// Crit logs a message with the Critical severity.
func (l *Log) Crit(msg string) { l.logmsg(LogCritical, msg) }

// Alert logs a message with the ALERT severity.
func (l *Log) Alert(msg string) { l.logmsg(LogAlert, msg) }

// Emerg logs a message with the EMERG severity.
func (l *Log) Emerg(msg string) { l.logmsg(LogEmergency, msg) }

// Fatal is just a wrapper for standard lib log.Fatal function.
func (l *Log) Fatal(msg string) { log.Fatal(msg) }

// Panic is just a wrapper for standard lib log.Panic function.
func (l *Log) Panic(msg string) { log.Panic(msg) }

// Debugf logs a custom formatted message with the DEBUG severity (just like the fmt.Printf).
func (l *Log) Debugf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogDebug, v...) }

// Infof logs a custom formatted message with the INFO severity (just like the fmt.Printf).
func (l *Log) Infof(frmt string, v ...interface{}) { l.logmsgf(frmt, LogInfo, v...) }

// Noticef logs a custom formatted message with the NOTIC severity (just like the fmt.Printf).
func (l *Log) Noticef(frmt string, v ...interface{}) { l.logmsgf(frmt, LogNotice, v...) }

// Warningf logs a custom formatted message with the WARN severity (just like the fmt.Printf).
func (l *Log) Warningf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogWarning, v...) }

// Errf logs a custom formatted message with the ERROR severity (just like the fmt.Printf).
func (l *Log) Errf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogError, v...) }

// Critf logs a custom formatted message with the CRIT severity (just like the fmt.Printf).
func (l *Log) Critf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogCritical, v...) }

// Alertf logs a custom formatted message with the ALERT severity (just like the fmt.Printf).
func (l *Log) Alertf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogAlert, v...) }

// Emergf logs a custom formatted message with the EMERG severity (just like the fmt.Printf).
func (l *Log) Emergf(frmt string, v ...interface{}) { l.logmsgf(frmt, LogEmergency, v...) }

// Fatalf is just a wrapper for standard lib log.Fatal function.
func (l *Log) Fatalf(frmt string, v ...interface{}) { log.Fatalf(frmt, v...) }

// Panicf is just a wrapper for standard lib log.Panicf function.
func (l *Log) Panicf(frmt string, v ...interface{}) { log.Panicf(frmt, v...) }

/*
// Printf is a generic logging method derived directly from standard log package and adapted for this advanced Log type.
func (l *Log) Printf(format string, v...interface{}) {
    for _, lg := range l.loggers {
        lg.Printf(format, v)
    }
}

// Print is a generic logging method derived directly from standard log package and adapted for this advanced Log type.
func (l *Log) Print(v...interface{}) {
    for _, lg := range l.loggers {
        lg.Print(v)
    }
}

// Println is a generic logging method derived directly from standard log package and adapted for this advanced Log type.
func (l *Log) Println(v...interface{}) {
    for _, lg := range l.loggers {
        lg.Println(v)
    }
}
*/

// Priority is an enum type abstracting the log message priority (severity).
// The type is basically a copy of the Priority type from STDLIB Priority from 'log/syslog' package. The STDLIB version is
// currently not built for Windows (why?).
type Priority int
const (
    // LogEmergency represents the emergency priority (severity) for logger
	LogEmergency Priority = iota
    // LogAlert represents the alert priority (severity) for logger
	LogAlert
    // LogCritical represents the critical priority (severity) for logger
	LogCritical
    // LogError represents the error priority (severity) for logger
	LogError
    // LogWarning represents the warning priority (severity) for logger
	LogWarning
    // LogNotice represents the notice priority (severity) for logger
	LogNotice
    // LogInfo represents the informational priority (severity) for logger
	LogInfo
    // LogDebug represents the debug priority (severity) for logger
	LogDebug
)

// String representation of the priority.
func (p Priority) String() string {
	switch p {
	case LogEmergency:
		return "Emergency"
	case LogAlert:
		return "Alert"
	case LogCritical:
		return "Critical"
	case LogError:
		return "Error"
	case LogWarning:
		return "Warning"
	case LogNotice:
		return "Notice"
	case LogInfo:
		return "Info"
	case LogDebug:
		return "Debug"
	}
	return ""
}

// Handler is type defining a format and severity for the particular Logger.
type Handler struct {

	// Message format for this particular log handler
	Format string

	// Priority for this particular log handler
	Priority Priority
}

// NewHandler creates a new log handler with given message format and priority.
func NewHandler(fmt string, prio Priority) *Handler { return &Handler{fmt, prio} }
