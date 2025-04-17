package slog

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

type StationFormatter struct {
	Fields []string
}

func (f *StationFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var keys = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		if k != "prefix" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	prefixFieldClashes(entry.Data)
	b := &bytes.Buffer{}

	f.printColored(b, entry, keys)

	return b.Bytes(), nil
}

func (f *StationFormatter) printColored(b *bytes.Buffer, entry *logrus.Entry, keys []string) {
	var levelColor string

	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = ansi.Green
	case logrus.WarnLevel:
		levelColor = ansi.Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = ansi.Red
	default:
		levelColor = ansi.Blue
	}
	levelText := "[" + strings.ToUpper(entry.Level.String()) + "]"
	_, _ = fmt.Fprintf(b, "%s %+5s %s \n", entry.Time.Format("2006-01-02 15:04:05"), levelText, entry.Message)
	for _, key := range keys {
		v := entry.Data[key]
		if f.Fields == nil || len(f.Fields) == 0 || fieldsInArray(key, f.Fields) {
			_, _ = fmt.Fprintf(b, " %s=%+v", key, v)
		}
	}
}

func fieldsInArray(key string, fields []string) bool {
	for _, field := range fields {
		if field == key {
			return true
		}
	}
	return false
}

func prefixFieldClashes(data logrus.Fields) {
	if _, ok := data["time"]; ok {
		data["fields.time"] = data["time"]
	}
	if _, ok := data["msg"]; ok {
		data["fields.msg"] = data["msg"]
	}
	if _, ok := data["level"]; ok {
		data["fields.level"] = data["level"]
	}
}
