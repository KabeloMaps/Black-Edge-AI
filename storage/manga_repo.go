package storage

import (
	"context"
	"errors"
	"time"

	"blackedge-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertManga(ctx context.Context, m models.NormalizedManga) error {

	if MangaCollection == nil {
		return errors.New("mongo not initialized")
	}

	filter := bson.M{"hash": m.Hash}

	update := bson.M{
		"$set": bson.M{
			"title":       m.Title,
			"description": m.Description,
			"cover_url":   m.CoverURL,
			"source":      m.Source,
			"genres":      m.Genres,
			"url":         m.URL,
			"updated_at":  time.Now(),
		},
		"$setOnInsert": bson.M{
			"hash":       m.Hash,
			"created_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := MangaCollection.UpdateOne(ctx, filter, update, opts)
	return err
}