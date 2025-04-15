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

func Initialize(ctx context.Context, logger *StationLogger, writer *os.File) error {
	switch logger.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	switch logger.LogFormat {
	case "discard":
		logrus.SetOutput(io.Discard)
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		if logger.LogFormat == "stdout" {
			logrus.SetOutput(os.Stdout)
		} else {
			logrus.SetOutput(writer)
		}
		for _, v := range logger.SkipTextFields {
			t := strings.SplitN(v, "=", 2)
			fieldName := t[0]
			fieldValue := t[1]
			log.Skip(log.Field(fieldName), fieldValue)
		}
		logrus.SetFormatter(&StationFormatter{})
	}

	return nil
}

func New(ctx context.Context) (*logrus.Logger, error) {
	newLogger := logrus.New()
	return newLogger, nil
}
