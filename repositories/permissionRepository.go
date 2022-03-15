package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type PermissionRepository struct{}

func (pr *PermissionRepository) SavePermission(permission *models.Permission) (*models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	permissionModel := mogo.NewDoc(permission).(*models.Permission)
	err := mogo.Save(permissionModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return permissionModel, err
}

func (pr *PermissionRepository) SavePermissions(permissions []*models.Permission) ([]*models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	for p := range permissions {
		permissionModel := mogo.NewDoc(permissions[p]).(*models.Permission)
		err := mogo.Save(permissionModel)
		if vErr, ok := err.(*mogo.ValidationError); ok {
			return nil, vErr
		}
		permissions[p] = permissionModel
	}
	return permissions, nil
}

func (pr *PermissionRepository) FindAll() ([]*models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	permissionDoc := mogo.NewDoc(models.Permission{}).(*models.Permission)
	var results []*models.Permission
	err := permissionDoc.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pr *PermissionRepository) FindById(id string) (*models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	permissionDoc := mogo.NewDoc(models.Permission{}).(*models.Permission)
	err := permissionDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, permissionDoc)
	if err != nil {
		return nil, err
	}
	return permissionDoc, nil
}
