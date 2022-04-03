package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type eventRepository struct{}

func (er *eventRepository) SaveEvent(event *models.Event) (*models.Event, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	eventModel := mogo.NewDoc(event).(*models.Event)
	err := mogo.Save(eventModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return eventModel, err
}

func (er *eventRepository) FindById(id string) (*models.Event, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	eventDoc := mogo.NewDoc(models.Event{}).(*models.Event)
	err := eventDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, eventDoc)
	if err != nil {
		return nil, err
	}
	return eventDoc, nil
}
