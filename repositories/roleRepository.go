package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type roleRepository struct{}

func (rr *roleRepository) SaveRole(role *models.Role) (*models.Role, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	roleModel := mogo.NewDoc(role).(*models.Role)
	err := mogo.Save(roleModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return roleModel, err
}

func (rr *roleRepository) FindByName(name string) (*models.Role, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	roleDoc := mogo.NewDoc(models.Role{}).(*models.Role)
	err := roleDoc.FindOne(bson.M{"name": name}, roleDoc)
	if err != nil {
		return nil, err
	}
	return roleDoc, nil
}

func (rr *roleRepository) FindById(id string, fetchRelationships bool) (interface{}, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	roleDoc := mogo.NewDoc(models.Role{}).(*models.Role)
	err := roleDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, roleDoc)
	if err != nil {
		return nil, err
	}
	if fetchRelationships {
		var roleResult models.RoleResult
		var permissions []models.Permission
		_ = roleDoc.Populate("Permissions").Find(bson.M{}).All(&permissions)
		_ = copier.Copy(&roleResult, roleDoc)
		roleResult.Permissions = permissions

		return &roleResult, nil
	}
	return roleDoc, nil
}

func (rr *roleRepository) FindAll() ([]*models.Role, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	roleDoc := mogo.NewDoc(models.Role{}).(*models.Role)
	var results []*models.Role
	err := roleDoc.Find(nil).Q().Sort("-_created").All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
