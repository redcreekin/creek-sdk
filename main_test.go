package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

func setupRouter() *gin.Engine {
	engine := gin.Default()
	return engine
}

func isValidJson(loader gojsonschema.JSONLoader, source []byte) (bool, []string) {
	documentLoader := gojsonschema.NewBytesLoader(source)
	result, err := gojsonschema.Validate(loader, documentLoader)
	if err != nil {
		errs := []string{err.Error()}
		return false, errs
	}
	if !result.Valid() {
		errs := []string{}
		for _, desc := range result.Errors() {
			errs = append(errs, desc.String())
		}
		return false, errs
	}
	return true, []string{}
}
