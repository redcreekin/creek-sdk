package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"net/http/httptest"
	"testing"
)

func environmentSchemaSetup() (gojsonschema.JSONLoader, error) {
	envSchema := GetEnvironmentRequestJsonSchema()
	schemaBytes, err := json.Marshal(envSchema)
	if err != nil {
		return nil, err
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaBytes)
	return schemaLoader, nil
}

func TestEnvironmentRequestSchema(t *testing.T) {
	engine := gin.Default()
	loader, err := environmentSchemaSetup()
	if err != nil {
		t.Fatal(err)
	}
	engine.POST("/environments", func(c *gin.Context) {
		var request EnvironmentRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			fmt.Println("Error binding JSON:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		environmentJson, err := json.Marshal(request)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		valid, errs := isValidJson(loader, environmentJson)
		if !valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Schema validation failed", "details": errs})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": request})
	})

	body := []byte(`{"environment_name": "DEV", "dynamic_infrastructure": false, "enable_guided_failure": false}`)
	req, _ := http.NewRequest("POST", "/environments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
