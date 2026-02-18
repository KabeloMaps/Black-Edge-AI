package storage

import (
	"blackedge-backend/database"
	"blackedge-backend/models"
	"context"
)

func SaveNormalizedManga(data []models.NormalizedManga) error {
	collection := database.MongoClient.Database("manga_engine").Collection("normalized_manga")

	var docs []interface{}
	for _, m := range data {
		docs = append(docs, m)
	}

	if len(docs) == 0 {
		return nil
	}

	_, err := collection.InsertMany(context.Background(), docs)
	return err
}
