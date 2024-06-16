package utils

import (
	"log"
	"os"
)

// LogError logs the provided error message if the error is not nil
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be defaultValue if the variable is not present.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// ValidateData checks if the provided data is valid
func ValidateData(data interface{}) bool {
	// Implement validation logic
	// Placeholder: return true as a default behavior
	return true
}
