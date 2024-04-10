package app

import (
	"os"
	"strconv"
)

func GetEnvString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		value, _ := strconv.Atoi(value)
		return value
	}
	return fallback
}
