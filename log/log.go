package slog

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/rockbears/log"
	"github.com/sirupsen/logrus"
)

const (
	HeaderRequestID      = "X-Request-ID"
	ExtraFieldSignature  = "Signature"
	ExtraFieldLine       = "Line"
	ExtraFieldTerminated = "Terminated"
)

type Conf struct {
	Level          string
	Format         string
	Options        []string
	TextFields     []string
	SkipTextFields []string

	SyslogHost     string
	SyslogPort     string
	SyslogProtocol string
	SyslogTag      string

	EventName string
	EventID   string
}

func Initialize(ctx context.Context, logger *Conf, writer *os.File) *log.Logger {
	loggersObject := logrus.New()
	switch logger.Level {
	case "debug":
		loggersObject.SetLevel(logrus.DebugLevel)
	case "info":
		loggersObject.SetLevel(logrus.InfoLevel)
	case "warn":
		loggersObject.SetLevel(logrus.WarnLevel)
	case "error":
		loggersObject.SetLevel(logrus.ErrorLevel)
	default:
		loggersObject.SetLevel(logrus.InfoLevel)
	}

	switch logger.Format {
	case "discard":
		loggersObject.SetOutput(io.Discard)
	case "json":
		loggersObject.SetFormatter(&logrus.JSONFormatter{})
	default:
		if logger.Format == "stdout" {
			loggersObject.SetOutput(os.Stdout)
		} else {
			loggersObject.SetOutput(writer)
		}
		loggersObject.SetFormatter(&StationFormatter{
			Fields: logger.TextFields,
		})
	}
	obj := log.NewWithFactory(log.NewLogrusWrapper(loggersObject))
	for _, v := range logger.SkipTextFields {
		t := strings.SplitN(v, "=", 2)
		fieldName := t[0]
		fieldValue := t[1]
		obj.UnregisterField(log.Field(fieldName))
		obj.Skip(log.Field(fieldName), fieldValue)
	}
	return obj
}
