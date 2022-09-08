package log

import (
	"github.com/sirupsen/logrus"
	"github.com/chocogem/slackbot-golang/pkg/config"
)

type LogLevel string
const (
	INFO  LogLevel = "INFO"
	DEBUG LogLevel = "DEBUG"
	TRACE LogLevel = "TRACE"
)

type LogrusImplement struct {
	logger *logrus.Logger
}
func NewLogrusLogger(cfg config.Config) *logrus.Logger {
	log := logrus.New()

	var setLogLevel logrus.Level
	var level = LogLevel(cfg.LogrusLogLevel)

	switch level {
	case DEBUG:
		setLogLevel = logrus.DebugLevel
	case TRACE:
		setLogLevel = logrus.TraceLevel
	default:
		setLogLevel = logrus.InfoLevel
	}

	log.Formatter = &logrus.JSONFormatter{}
	log.SetLevel(setLogLevel)

	return log
}
func NewImplementLogrusLogger(logger *logrus.Logger) *LogrusImplement {
	return &LogrusImplement{logger}
}

func (l *LogrusImplement) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *LogrusImplement) Error(msg string) {
	l.logger.Error(msg)
}

func (l *LogrusImplement) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *LogrusImplement) Info(msg string) {
	l.logger.Info(msg)
}
