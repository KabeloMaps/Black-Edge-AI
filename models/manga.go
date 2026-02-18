package models

// Raw scraped data from any source
type ScrapedManga struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	CoverImage  string   `json:"cover_image"`
	Genres      []string `json:"genres"`
	Language    string   `json:"language"`
	Status      string   `json:"status"`
	SourceURL   string   `json:"source_url"`
}

// Normalized data stored in Mongo
type NormalizedManga struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Slug        string    `json:"slug" bson:"slug"`
	Description string    `json:"description" bson:"description"`
	Author      string    `json:"author" bson:"author"`
	CoverImage  string    `json:"cover_image" bson:"cover_image"`
	Genres      []string  `json:"genres" bson:"genres"`
	Language    string    `json:"language" bson:"language"`
	Status      string    `json:"status" bson:"status"`
	Source      string    `json:"source" bson:"source"`
	SourceURL   string    `json:"source_url" bson:"source_url"`
	CreatedAt   int64     `json:"created_at" bson:"created_at"`
	UpdatedAt   int64     `json:"updated_at" bson:"updated_at"`
	Version     int       `json:"version" bson:"version"`
}

// API alias
type Manga = NormalizedManga
