package models

import "time"

type ScrapedManga struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CoverURL    string   `json:"cover_url"`
	URL         string   `json:"url"`
	Genres      []string `json:"genres"`
}

type NormalizedManga struct {
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	CoverURL    string    `bson:"cover_url" json:"cover_url"`
	URL         string    `bson:"url" json:"url"`
	Genres      []string  `bson:"genres" json:"genres"`
	Source      string    `bson:"source" json:"source"`
	Hash        string    `bson:"hash" json:"hash"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}