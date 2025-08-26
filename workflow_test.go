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

func workflowSchemaSetup() (gojsonschema.JSONLoader, error) {
	envSchema := GetWorkflowRequestJsonSchema()
	schemaBytes, err := json.Marshal(envSchema)
	if err != nil {
		return nil, err
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaBytes)
	return schemaLoader, nil
}

func TestWorkflowRequestSchema(t *testing.T) {
	engine := gin.Default()
	loader, err := workflowSchemaSetup()
	if err != nil {
		t.Fatal(err)
	}
	engine.POST("/project_groups", func(c *gin.Context) {
		var request ProjectGroupRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		environmentJson, err := json.Marshal(request)
		if err != nil {
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

	body := []byte(`{"station.workflow.name": "Test Workflow", "project_type": "deploy"}`)
	req, _ := http.NewRequest("POST", "/project_groups", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWorkflowResponseOutput(t *testing.T) {

}
