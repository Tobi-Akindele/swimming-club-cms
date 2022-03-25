package utils

import (
	"github.com/goonode/mogo"
	"os"
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

func ExtractRefFieldId(refField mogo.RefField) string {
	return refField.ID.Hex()
}

func MapContainsKey(iMap map[string]string, key string) bool {
	_, ok := iMap[key]
	return ok
}
