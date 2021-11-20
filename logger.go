package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type FormatterType int64

const (
	JSON FormatterType = 0
	TEXT               = 1
)

type Log struct {
	Title   string
	Reason  string
	Payload string
	Topic   string
	Message string
}

func DoLog(title string, reason string, payload string, topic string, message string) {
	l := Log{
		Title:   title,
		Reason:  reason,
		Payload: payload,
		Topic:   topic,
		Message: message,
	}
	l.log()
}

type LogConfig struct {
	Format   FormatterType
	LogLevel log.Level
}

var logConfig LogConfig

func NewLogConfig(format FormatterType, level log.Level) {
	logConfig.LogLevel = level
	logConfig.Format = format
}

func init() {
	if logConfig.Format == JSON {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(logConfig.LogLevel)
}

func (l Log) Error() {
	l.entry().Error(l.Message)
}

func (l Log) Info() {
	l.entry().Info(l.Message)
}

func (l Log) Warn() {
	l.entry().Warn(l.Message)
}

func (l Log) Panic() {
	l.entry().Panic(l.Message)
}

func (l Log) Fatal() {
	l.entry().Fatal(l.Message)
}

func (l Log) Debug() {
	l.entry().Debug(l.Message)
}

func (l Log) Trace() {
	l.entry().Trace(l.Message)
}

func (l Log) entry() *log.Entry {
	return log.WithFields(log.Fields{
		"title":   l.Title,
		"reason":  l.Reason,
		"payload": l.Payload,
		"topic":   l.Topic,
	})
}

func (l Log) log() {
	if logConfig.LogLevel == log.InfoLevel {
		l.Info()
	} else if logConfig.LogLevel == log.DebugLevel {
		l.Debug()
	} else if logConfig.LogLevel == log.TraceLevel {
		l.Trace()
	} else if logConfig.LogLevel == log.FatalLevel {
		l.Fatal()
	} else if logConfig.LogLevel == log.PanicLevel {
		l.Panic()
	} else if logConfig.LogLevel == log.ErrorLevel {
		l.Error()
	} else if logConfig.LogLevel == log.WarnLevel {
		l.Warn()
	}
}
