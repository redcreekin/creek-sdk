package hook

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ContextHook struct {
	Keys []string
	Ctx  context.Context
}

func (h *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *ContextHook) Fire(entry *logrus.Entry) error {
	for _, key := range h.Keys {
		if value, ok := h.Ctx.Value(key).(string); ok {
			entry.Data[key] = value
		}
	}
	return nil
}
