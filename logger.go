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
	Title string
	Reason string
	Payload string
	Topic string
	Message string
}


type LogConfig struct {
	Format FormatterType
	LogLevel log.Level
}

var logConfig LogConfig

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
		"title": l.Title,
		"reason": l.Reason,
		"payload": l.Payload,
		"topic": l.Topic,
	})
}

