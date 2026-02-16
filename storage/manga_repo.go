package storage

import (
	"blackedge-backend/database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveRawMangaData(mangaData interface{}) error {
	
	collection := database.MongoClient.Database("blackedge").Collection("raw_manga")
	_, err := collection.InsertOne(context.Background(), bson.M{"data": mangaData})
	return err
}