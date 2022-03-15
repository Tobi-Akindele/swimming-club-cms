package utils

import (
	"github.com/goonode/mogo"
	"os"
)

const (
	APP_NAME            = "APP_NAME"
	PORT                = "PORT"
	MONGO_DB_HOST       = "MONGO_DB_HOST"
	DB                  = "DB"
	DOB_DATE_FORMAT     = "02-01-2006"
	JWT_SECRET_KEY      = "JWT_SECRET_KEY"
	JWT_TOKEN_EXPIRY    = "JWT_TOKEN_EXPIRY"
	AUTHORIZATION       = "Authorization"
	USER                = "user"
	SUPER_ADMIN         = "SUPER ADMIN"
	SUPER_USER_PASS_KEY = "SUPER_USER_PASS_KEY"
	USER_PERMISSIONS    = "USER_PERMISSIONS"
)

func GetEnv(param string, defaultValue string) string {
	envValue := os.Getenv(param)
	if len(envValue) == 0 {
		return defaultValue
	}
	return envValue
}

func ConvertRefFieldSliceToStringMap(refFieldSlice mogo.RefFieldSlice) map[string]string {
	result := map[string]string{}
	for refIdx := range refFieldSlice {
		result[refFieldSlice[refIdx].ID.String()] = refFieldSlice[refIdx].ID.String()
	}
	return result
}

func ConvertRefFieldSliceToStringSlice(refFieldSlice mogo.RefFieldSlice) []string {
	result := make([]string, len(refFieldSlice))
	for refIdx := range refFieldSlice {
		result[refIdx] = refFieldSlice[refIdx].ID.Hex()
	}
	return result
}

func HasPermission(permissions map[string]string, permission string) bool {
	_, ok := permissions[permission]
	return ok
}
