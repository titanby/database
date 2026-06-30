package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(ctx context.Context, collection *mongo.Collection, filter interface{}, result interface{}, opts ...*options.FindOptions) error {
	ctx = dbContext(ctx)
	cur, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func UpdateOne(ctx context.Context, collection *mongo.Collection, filter, update interface{}, opts ...*options.UpdateOptions) error {
	ctx = dbContext(ctx)
	_, err := collection.UpdateOne(ctx, filter, update, opts...)
	return err
}
