package database

// func updateOne(collection *mongo.Collection, filter, update interface{}) error {
// 	_, err := collection.UpdateOne(Ctx(), filter, update, options.Update().SetUpsert(true))
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Update", update,
// 			"Collection", collection.Name(),
// 			"Database", collection.Database().Name(),
// 		).Error("Error updating data")
// 	}
// 	return err
// }
//
// func UpdateOneLog(colName string, filter, update interface{}) error {
// 	err := updateOne(logDatabase().Collection(colName), filter, bson.M{"$set": update})
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Update", update,
// 			"Collection", colName,
// 		).Error("Error updating data")
// 	}
// 	return err
// }
//
// func UpdateOne(colName string, filter, update interface{}) error {
// 	err := updateOne(Collection(colName), filter, bson.M{"$set": update})
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Update", update,
// 			"Collection", colName,
// 		).Error("Error updating data")
// 	}
// 	return err
// }
//
// func UpdateOneLogAndUnsetError(colName string, filter, update interface{}) error {
// 	updater := bson.M{"$set": update}
// 	updater["$unset"] = bson.M{"sender_error": nil}
// 	return updateOne(logDatabase().Collection(colName), filter, updater)
// }
//
// func UpdateOneAndUnsetError(colName string, filter, update interface{}) error {
// 	updater := bson.M{"$set": update}
// 	updater["$unset"] = bson.M{"sender_error": nil}
// 	return updateOne(Collection(colName), filter, updater)
// }
//
// func updateMany(colName string, filter, update interface{}) error {
// 	_, err := Collection(colName).UpdateMany(Ctx(), filter, update, options.Update().SetUpsert(true))
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Update", update,
// 			"Collection", colName,
// 		).Error("Error updating data")
// 	}
// 	return err
// }
//
// func UpdateMany(colName string, filter, update interface{}) error {
// 	return updateMany(colName, filter, bson.M{"$set": update})
// }
//
// func UpdateManyAndUnsetError(colName string, filter interface{}, update interface{}) error {
// 	updater := bson.M{"$set": update}
// 	updater["$unset"] = bson.M{"sender_error": nil}
// 	return updateMany(colName, filter, updater)
// }
//
// func DeleteOne(name string, filter bson.M) error {
// 	_, err := Collection(name).DeleteOne(Ctx(), filter)
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Collection", name,
// 			"Filter", filter,
// 		).Error("Error deleting data")
// 	}
// 	return err
// }
//
// func DeleteMany(colName string, filter interface{}) error {
// 	_, err := Collection(colName).DeleteMany(Ctx(), filter)
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Collection", colName,
// 			"Filter", filter,
// 		).Error("Error deleting data")
// 	}
// 	return err
// }
//
// func DeleteAll(colName string) error {
// 	if _, err := Collection(colName).DeleteMany(Ctx(), bson.M{}); err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Collection", colName,
// 		).Error("Error deleting data")
// 		return err
// 	} else {
// 		app.LogWith(
// 			"Collection", colName,
// 		).Debug("Data deleted")
// 	}
// 	return nil
// }
//
// func findOne(collection *mongo.Collection, filter, sortBy interface{}, result interface{}) error {
// 	opt := options.FindOne().SetSort(sortBy)
// 	return findOneWithOptions(collection, filter, result, opt)
// }
//
// func FindOneInSecondary(collectionName string, filter, sortBy interface{}, result interface{}) error {
// 	return findOne(SecondaryCollection(collectionName), filter, sortBy, result)
// }
//
// func FindOne(collectionName string, filter, sortBy interface{}, result interface{}) error {
// 	return findOne(Collection(collectionName), filter, sortBy, result)
// }
//
// func findOneWithOptions(collection *mongo.Collection, filter, result interface{}, opts *options.FindOneOptions) error {
// 	err := collection.FindOne(Ctx(), filter, opts).Decode(result)
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Collection", collection.Name(),
// 			"Options", opts,
// 		).Error("Error FindOne")
// 	}
// 	return err
// }
//
// func FindOneWithOptionsInSecondary(collectionName string, filter, result interface{}, opts *options.FindOneOptions) error {
// 	return findOneWithOptions(SecondaryCollection(collectionName), filter, result, opts)
// }
//
// func FindOneWithOptions(collectionName string, filter, result interface{}, opts *options.FindOneOptions) error {
// 	return findOneWithOptions(Collection(collectionName), filter, result, opts)
// }
//
// func find(collection *mongo.Collection, filter, sortBy interface{}, limit int64, result interface{}) error {
// 	cur, err := collection.Find(Ctx(), filter, options.Find().SetSort(sortBy).SetLimit(limit))
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Collection", collection.Name(),
// 		).Error("Error getting data")
// 		return err
// 	}
// 	defer cur.Close(Ctx())
//
// 	if err = cur.All(Ctx(), result); err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Filter", filter,
// 			"Collection", collection.Name(),
// 		).Error("Error getting data")
// 		return err
// 	}
// 	return nil
// }
//
// func FindInSecondary(collectionName string, filter, sortBy interface{}, limit int64, result interface{}) error {
// 	return find(SecondaryCollection(collectionName), filter, sortBy, limit, result)
// }
//
// func Find(collectionName string, filter, sortBy interface{}, limit int64, result interface{}) error {
// 	return find(Collection(collectionName), filter, sortBy, limit, result)
// }
//
// func GetLastDocument(name string, field string, res interface{}) error {
// 	opts := options.FindOne().SetSort(bson.M{field: -1})
// 	err := Collection(name).FindOne(Ctx(), bson.M{}, opts).Decode(res)
// 	if err != nil {
// 		app.LogWith(
// 			"Error", err,
// 			"Collection", name,
// 			"Field", field,
// 		).Error("Error getting last document")
// 	}
// 	return err
// }
