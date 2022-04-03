package db

import (
	"github.com/goonode/mogo"
	"swimming-club-cms-be/utils"
)

var mongoDBConnection *mogo.Connection = nil

func GetConnection() *mogo.Connection {
	if mongoDBConnection != nil {
		return mongoDBConnection
	}
	config := &mogo.Config{
		ConnectionString: utils.GetEnv(utils.MONGO_DB_HOST, ""),
		Database:         utils.GetEnv(utils.DB, ""),
	}
	mongoDBConnection, err := mogo.Connect(config)
	if err != nil {
		panic(err)
	}
	mongoDBConnection.Session.SetSyncTimeout(0)
	return mongoDBConnection
}

func CloseConnection(conn *mogo.Connection) {
	conn.Session.Close()
}
