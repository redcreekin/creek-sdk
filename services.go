package sdk

import (
	"database/sql/driver"
	"encoding/json"
)

const (
	TypeRepository = "repository"
	TypeGRPC       = "grpc"
	TypeAPI        = "api"
	TypeUI         = "ui"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}
	return json.Unmarshal(value.([]byte), j)
}
