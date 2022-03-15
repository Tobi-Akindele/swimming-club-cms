package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type RoleRepository struct{}

func (rr *RoleRepository) SaveRole(role *models.Role) (*models.Role, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	roleModel := mogo.NewDoc(role).(*models.Role)
	err := mogo.Save(roleModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return roleModel, err
}

func (rr *RoleRepository) FindByName(name string) (*models.Role, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	roleDoc := mogo.NewDoc(models.Role{}).(*models.Role)
	err := roleDoc.FindOne(bson.M{"name": name}, roleDoc)
	if err != nil {
		return nil, err
	}
	return roleDoc, nil
}

func (rr *RoleRepository) FindById(id string) (*models.Role, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	roleDoc := mogo.NewDoc(models.Role{}).(*models.Role)
	err := roleDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, roleDoc)
	if err != nil {
		return nil, err
	}
	return roleDoc, nil
}
