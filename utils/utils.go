package utils

import (
	"github.com/gin-gonic/gin"
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

func ConvertRefFieldSliceToMap(refFieldSlice mogo.RefFieldSlice) map[string]int {
	result := make(map[string]int, len(refFieldSlice))
	for k, refField := range refFieldSlice {
		result[refField.ID.Hex()] = k
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

func MapContainsKey1(iMap map[string]int, key string) bool {
	_, ok := iMap[key]
	return ok
}

func ConvertPermissionSliceToMap(permissionsSrc []models.Permission) map[string]string {
	permissions := make(map[string]string, len(permissionsSrc))
	for _, permission := range permissionsSrc {
		permissions[permission.Value] = permission.Name
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

func IsAdmin(ctx *gin.Context) (*models.User, bool) {
	if rawUser, ok := ctx.Get(USER); ok {
		user, _ := rawUser.(*models.User)
		return user, user.Admin
	}
	//Ideally this line is unreachable
	return nil, false
}
