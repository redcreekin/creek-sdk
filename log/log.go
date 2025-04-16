package slog

import (
	"context"
	"github.com/rockbears/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const (
	HeaderRequestID      = "X-Request-ID"
	ExtraFieldSignature  = "Signature"
	ExtraFieldLine       = "Line"
	ExtraFieldTerminated = "Terminated"
)

type StationLogger struct {
	Level       string   `toml:"level" validate:"required,oneof=debug info warn error fatal panic" comment:"Level of logs need to be added in file" default:"info"` // Log level
	LogPath     string   `toml:"logPath" validate:"required,dirpath" comment:"Location where logfile will be created" default:""`                                   // Log file path
	LogFileName string   `toml:"logFileName" validate:"required" default:"" comment:"Log file name"`                                                                // Log file name
	LogFormat   string   `toml:"logFormat" validate:"required,oneof=discard stdout json file" comment:"Log format" default:"file"`                                  // Log format
	LogOptions  []string `toml:"logOptions" validate:"omitempty" comment:"Log options"`                                                                             // Log options

	SkipTextFields []string `toml:"skipTextFields" validate:"omitempty" comment:"Skip text fields"` // Skip text fields
}

func Initialize(ctx context.Context, logger *StationLogger, writer *os.File) *log.Logger {
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

	switch logger.LogFormat {
	case "discard":
		loggersObject.SetOutput(io.Discard)
	case "json":
		loggersObject.SetFormatter(&logrus.JSONFormatter{})
	default:
		if logger.LogFormat == "stdout" {
			loggersObject.SetOutput(os.Stdout)
		} else {
			loggersObject.SetOutput(writer)
		}
		loggersObject.SetFormatter(&StationFormatter{})
	}
	obj := log.NewWithFactory(log.NewLogrusWrapper(loggersObject))
	for _, v := range logger.SkipTextFields {
		t := strings.SplitN(v, "=", 2)
		fieldName := t[0]
		fieldValue := t[1]
		obj.Skip(log.Field(fieldName), fieldValue)
	}
	return obj
}
