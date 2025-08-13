package slog

import (
	"context"
	"fmt"
	"github.com/rockbears/log"
	"os"
	"path/filepath"
	"testing"
)

func TestSkipLoggingFields(t *testing.T) {
	// Test logging functionality here
	ctx := context.Background()
	tempDir := os.TempDir()
	logFile := filepath.Join(tempDir, "project.log")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	logConf := Conf{
		Level:          "info",
		Format:         "stdout",
		TextFields:     []string{},
		SkipTextFields: []string{"caller=%", "source_file=%", "source_line=%"},
	}

	logger := Initialize(ctx, &logConf, file)
	logger.Info(ctx, "This is Info message")

	t.Logf("Logged Info message")

}

func TestAddLoggingFields(t *testing.T) {
	ctx := context.WithValue(context.Background(), Component, "TestComponent")
	tempDir := os.TempDir()
	logFile := filepath.Join(tempDir, "project.log")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	logConf := Conf{
		Level:          "info",
		Format:         "stdout",
		TextFields:     []string{"component", "action", "consumer"},
		SkipTextFields: []string{"caller=%", "source_file=%", "source_line=%"},
	}
	logger := Initialize(ctx, &logConf, file)

	logger.RegisterField(log.Field("application"))

	newCtx := context.WithValue(ctx, Action, "DemoApp")

	logger.Info(newCtx, "This is Info message -- with Component Field")
	t.Logf("Checks if Logging field is added with value")
}
