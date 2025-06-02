package sdk

import (
	"context"
	"encoding/json"

	"github.com/rockbears/log"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

var schema *jsonschema.Schema

func ValidateSchema(ctx context.Context, requestData interface{}, logger *log.Logger) error {
	data, err := json.Marshal(requestData)
	if err != nil {
		logger.Debug(ctx, "Failed to marshal request data %s", err)
		return err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Debug(ctx, "Failed to unmarshal request data %s", err)
		return err
	}
	err = schema.Validate(result)
	if err != nil {
		logger.Debug(ctx, "Request data validation failed: %s", err)
		return err
	}
	logger.Debug(ctx, "Request data validation succeeded")
	return nil
}
