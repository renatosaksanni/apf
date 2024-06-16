package utils

import (
	"os"
	"testing"
)

func TestLogError(t *testing.T) {
	LogError(nil) // Should not produce any log
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV", "test_value")
	if val := GetEnv("TEST_ENV", "default"); val != "test_value" {
		t.Fatalf("Expected 'test_value', got %s", val)
	}

	if val := GetEnv("NON_EXISTENT_ENV", "default"); val != "default" {
		t.Fatalf("Expected 'default', got %s", val)
	}
}
