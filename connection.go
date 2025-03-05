package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	connection struct {
		dbName          string
		primaryClient   *mongo.Client // master
		secondaryClient *mongo.Client // slave
	}
	Connection interface {
		PrimaryClient() *mongo.Client
		PrimaryDatabase() *mongo.Database
		SecondaryClient() *mongo.Client
		SecondaryDatabase() *mongo.Database
	}
)

func (conn *connection) PrimaryClient() *mongo.Client {
	return conn.primaryClient
}

func (conn *connection) PrimaryDatabase() *mongo.Database {
	return conn.primaryClient.Database(conn.dbName)
}

func (conn *connection) SecondaryClient() *mongo.Client {
	return conn.secondaryClient
}

func (conn *connection) SecondaryDatabase() *mongo.Database {
	return conn.secondaryClient.Database(conn.dbName)
}
