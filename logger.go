package logandgo

import (
	"log"
	"os"
)

type LogLevel int64

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
	panicLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	panicLogger = log.New(os.Stderr, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
}

const (
	PANIC LogLevel = 0
	ERROR          = 1
	WARN           = 2
	INFO           = 3
	DEBUG          = 4
)

type Format int64

const (
	JSON Format = 0
	TEXT        = 1
)

type LogFields map[string]interface{}

func (lf LogFields) toJson() string {
	return toJson(lf)
}

func (lf LogFields) toText() string {
	return toText("", lf)
}

type Logger interface {
	Debug(fields LogFields)
	Info(fields LogFields)
	Warn(fields LogFields)
	Error(fields LogFields)
	Panic(fields LogFields)
}

func NewLogger(format Format, level int) Logger {
	switch format {
	case JSON:
		return NewJsonLogger(level)
	case TEXT:
		return NewTextLogger(level)
	default:
		panic("logger not defined")
	}
}

type JsonLogger struct {
	logLevel LogLevel
}

func NewJsonLogger(level int) *JsonLogger {
	return &JsonLogger{logLevel: LogLevel(level)}
}

func (j *JsonLogger) Debug(fields LogFields) {
	output(debugLogger, DEBUG, j.logLevel, fields.toJson())
}

func (j *JsonLogger) Error(fields LogFields) {
	output(errorLogger, ERROR, j.logLevel, fields.toJson())
}

func (j *JsonLogger) Info(fields LogFields) {
	output(infoLogger, INFO, j.logLevel, fields.toJson())
}

func (j *JsonLogger) Warn(fields LogFields) {
	output(warningLogger, WARN, j.logLevel, fields.toJson())
}

func (j *JsonLogger) Panic(fields LogFields) {
	output(panicLogger, PANIC, j.logLevel, fields.toJson())
}

type TextLogger struct {
	logLevel LogLevel
}

func NewTextLogger(level int) *TextLogger {
	return &TextLogger{logLevel: LogLevel(level)}
}

func (t *TextLogger) Debug(fields LogFields) {
	output(debugLogger, DEBUG, t.logLevel, fields.toText())
}

func (t *TextLogger) Error(fields LogFields) {
	output(errorLogger, ERROR, t.logLevel, fields.toText())
}

func (t *TextLogger) Info(fields LogFields) {
	output(infoLogger, INFO, t.logLevel, fields.toText())
}

func (t *TextLogger) Warn(fields LogFields) {
	output(warningLogger, WARN, t.logLevel, fields.toText())
}

func (t *TextLogger) Panic(fields LogFields) {
	output(panicLogger, PANIC, t.logLevel, fields.toText())
}
