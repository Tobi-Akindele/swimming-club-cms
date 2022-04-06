package utils

import (
	"github.com/goonode/mogo"
	"os"
	"swimming-club-cms-be/models"
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
		result[refFieldSlice[refIdx].ID.Hex()] = refFieldSlice[refIdx].ID.Hex()
	}
	return result
}

func RemoveRefFromRefSlice(refSlice mogo.RefFieldSlice, refIds []string) mogo.RefFieldSlice {
	result := mogo.RefFieldSlice{}
	for _, refField := range refSlice {
		if !contains(refIds, refField.ID.Hex()) {
			result = append(result, &mogo.RefField{ID: refField.ID})
		}
	}
	return result
}

func contains(elements []string, str string) bool {
	for _, s := range elements {
		if str == s {
			return true
		}
	}
	return false
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

func ConvertPermissionSliceToMap(permissionsSrc []models.Permission) map[string]string {
	permissions := map[string]string{}
	for i := range permissionsSrc {
		permissions[permissionsSrc[i].Value] = permissionsSrc[i].Name
	}
	return permissions
}

func ExtractUserIdsFromUserStructs(users []models.UserResult) []string {
	var result []string
	for _, user := range users {
		result = append(result, user.ID.Hex())
	}
	return result
}
