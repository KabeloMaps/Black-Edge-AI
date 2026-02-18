package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureIndexes() error {
	_, err := MangaCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"slug": 1},
		},
		{
			Keys: bson.M{"genres": 1},
		},
		{
			Keys: bson.M{"language": 1},
		},
		{
			Keys: bson.M{"status": 1},
		},
		{
			Keys: bson.M{"source": 1},
		},
	})
	return err
}
