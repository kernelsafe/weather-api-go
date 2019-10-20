package util

import "os"

// GetEnv returns the value for the key, If the key doesn't exist, returns defValue.
func GetEnv(key, defValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defValue
	}
	return value
}
