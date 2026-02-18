package ingestion// =========================
// RAW SCRAPED STRUCT (SOURCE LAYER)
// =========================

type ScrapedManga struct {
	Title       string   `json:"title"`        // Raw title from site
	Description string   `json:"description"`  // Raw description
	Author      string   `json:"author"`       // Raw author
	CoverImage  string   `json:"cover_image"`  // Image URL
	Genres      []string `json:"genres"`       // Raw genres
	Language    string   `json:"language"`     // Language code or label
	Status      string   `json:"status"`       // ongoing/completed/unknown
	SourceURL   string   `json:"source_url"`   // Original content URL
	SourceName  string   `json:"source_name"`  // Scraper ID (site name)
}

// =========================
// NORMALIZED STRUCT (CORE DOMAIN MODEL)
// =========================

type NormalizedManga struct {
	ID          string   `json:"id" bson:"_id"`         // Deterministic hash ID
	Title       string   `json:"title" bson:"title"`   // Normalized title
	Slug        string   `json:"slug" bson:"slug"`     // URL-safe slug
	Description string   `json:"description" bson:"description"`
	Author      string   `json:"author" bson:"author"`
	CoverImage  string   `json:"cover_image" bson:"cover_image"`
	Genres      []string `json:"genres" bson:"genres"`
	Language    string   `json:"language" bson:"language"`
	Status      string   `json:"status" bson:"status"` // ongoing/completed/unknown
	Source      string   `json:"source" bson:"source"` // scraper name
	SourceURL   string   `json:"source_url" bson:"source_url"`

	// System fields
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
	Version   int   `json:"version" bson:"version"`
}

// =========================
// API RESPONSE MODEL
// =========================

type MangaResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	CoverImage  string   `json:"cover_image"`
	Genres      []string `json:"genres"`
	Language    string   `json:"language"`
	Status      string   `json:"status"`
	Source      string   `json:"source"`
	SourceURL   string   `json:"source_url"`
}

// =========================
// PAGINATION MODEL
// =========================

type PaginatedMangaResponse struct {
	Data       []MangaResponse `json:"data"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	Total      int64           `json:"total"`
	TotalPages int             `json:"total_pages"`
}

