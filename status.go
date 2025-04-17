package sdk

import (
	"fmt"
	"time"
)

const (
	MonitoringStatusOK       = "OK"
	MonitoringStatusWarning  = "WARNING"
	MonitoringStatusCritical = "CRITICAL"
	MonitoringStatusUnknown  = "UNKNOWN"
)

type MonitoringStatus struct {
	Now   time.Time              `json:"now"`
	Lines []MonitoringStatusLine `json:"lines"`

	ServiceType     string `json:"-"`
	ServiceName     string `json:"-"`
	ServiceHostname string `json:"-"`
}

type MonitoringStatusLine struct {
	Status     string `json:"status"`
	Component  string `json:"component"`
	Value      string `json:"value"`
	Type       string `json:"type"`
	Service    string `json:"service"`
	Hostname   string `json:"hostname"`
	SessionID  string `json:"session,omitempty"`
	ConsumerID string `json:"consumer,omitempty"`
}

func (m *MonitoringStatusLine) GetStatus() string {
	return fmt.Sprintf("%s - %s: %s", m.Status, m.Component, m.Value)
}
