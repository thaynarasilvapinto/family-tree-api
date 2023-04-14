package logger

import (
	"fmt"
	"log"
	"os"
)

type myLogger struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
}

var logger *myLogger

func init() {
	logger = &myLogger{
		errorLog:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warningLog: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *myLogger) Error(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.errorLog.Output(2, msg)
}

func (l *myLogger) Warning(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.warningLog.Output(2, msg)
}

func (l *myLogger) Info(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.infoLog.Output(2, msg)
}

func Error(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	logger.errorLog.Output(2, msg)
}

func Warning(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	logger.warningLog.Output(2, msg)
}

func Info(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	logger.infoLog.Output(2, msg)
}
