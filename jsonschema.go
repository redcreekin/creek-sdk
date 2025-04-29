package sdk

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/invopop/jsonschema"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}
	return json.Unmarshal(value.([]byte), j)
}

func GetActionJsonSchema(actionModel interface{}) *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	actionSchema := reflector.Reflect(actionModel)
	return actionSchema
}
