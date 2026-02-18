package storage

import (
	"blackedge-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertManga(ctx context.Context, manga models.NormalizedManga) error {
	_, err := MangaCollection.InsertOne(ctx, manga)
	return err
}

// Batch insert (used by ingestion pipeline)
func InsertMangaBatch(ctx context.Context, mangaList []models.NormalizedManga) error {
	if len(mangaList) == 0 {
		return nil
	}

	docs := make([]interface{}, 0, len(mangaList))
	for _, m := range mangaList {
		docs = append(docs, m)
	}

	_, err := MangaCollection.InsertMany(ctx, docs)
	return err
}

// Production-grade upsert (recommended)
func UpsertManga(ctx context.Context, manga models.NormalizedManga) error {
	filter := bson.M{"id": manga.ID}
	update := bson.M{"$set": manga}
	opts := options.Update().SetUpsert(true)

	_, err := MangaCollection.UpdateOne(ctx, filter, update, opts)
	return err
}


