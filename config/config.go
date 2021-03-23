package config

import (
	"os"
	"strconv"
)

func GetEnvString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

func GetEnvInt(key string, defaultVal int) int {

	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}

	return intValue
}
