package ingestion

type SourceAdapter interface {
	Name() string
	BaseURL() string
	GetMangaList() ([]ScrapedManga, error)
}
