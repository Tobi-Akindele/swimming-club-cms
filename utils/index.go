package utils

import "os"

const (
	APP_NAME = "APP_NAME"
	PORT     = "PORT"
)

func GetEnv(param string, defaultValue string) string {
	envValue := os.Getenv(param)
	if len(envValue) == 0 {
		return defaultValue
	}
	return envValue
}
