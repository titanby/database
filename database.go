package database

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbPool = make(map[string]*connection)

func Connect(alias, dbName, connString string) error {
	if connString == "" || dbName == "" || alias == "" {
		return errors.New("empty one of connections param")
	}

	_, ok := dbPool[alias]
	if ok {
		return errors.New("database already exists")
	}

	if connString[:10] != "mongodb://" {
		connString = "mongodb://" + connString
	}

	writer, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		return err
	}

	if err = writer.Ping(context.Background(), nil); err != nil {
		return err
	}

	reader := writer

	if strings.Contains(connString, "readPreference=primary") {
		if reader, err = mongo.Connect(context.Background(), options.Client().ApplyURI(strings.Replace(connString, "readPreference=primary", "readPreference=secondary", 1))); err != nil {
			return err
		}
	}

	dbPool[alias] = &connection{
		dbName:          dbName,
		primaryClient:   writer,
		secondaryClient: reader,
	}

	return nil
}

func Disconnect(alias string) {
	client, ok := dbPool[alias]
	if !ok {
		return
	}
	_ = client.primaryClient.Disconnect(context.Background())
	_ = client.secondaryClient.Disconnect(context.Background())
	delete(dbPool, alias)
}

func GetConnection(alias string) Connection {
	client, ok := dbPool[alias]
	if !ok {
		return nil
	}
	return client
}
