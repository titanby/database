package database

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClient struct {
	dbName string
	writer *mongo.Client
	reader *mongo.Client
}

var dbPool = make(map[string]*dbClient)

func Connect(alias, dbName, connString string) error {
	// cfg := struct {
	// 	Config config `json:"database"`
	// }{}
	// // Read config
	// if err := app.GetConfig(&cfg); err != nil {
	// 	app.LogWith(
	// 		"Error", err,
	// 	).Fatal("Read Database config")
	// }

	// app.LogWith("Config", cfg.Config).Debug("Database config")

	// app.LogInfo("Connecting to database...")
	// connString := cfg.Config.ConnectionString
	// dbName = cfg.Config.Name

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

	dbPool[alias] = &dbClient{
		dbName: dbName,
		writer: writer,
		reader: reader,
	}

	return nil
}

func Disconnect(alias string) {
	client, ok := dbPool[alias]
	if !ok {
		return
	}
	_ = client.writer.Disconnect(context.Background())
	_ = client.reader.Disconnect(context.Background())
	delete(dbPool, alias)
}
