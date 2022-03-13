package utils

import "os"

const (
	APP_NAME        = "APP_NAME"
	PORT            = "PORT"
	MONGO_DB_HOST   = "MONGO_DB_HOST"
	DB              = "DB"
	DOB_DATE_FORMAT = "02-01-2006"
)

func GetEnv(param string, defaultValue string) string {
	envValue := os.Getenv(param)
	if len(envValue) == 0 {
		return defaultValue
	}
	return envValue
}
