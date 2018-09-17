package util

import "os"

// GetenvRequired returns the environment variable key.
// If it is not defined, panics the application.
func GetenvRequired(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic("Variable `" + key + "` is required, please set it in `.env` file.")
	}

	return val
}
