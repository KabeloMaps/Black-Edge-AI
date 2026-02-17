package ingestion

type ScrapedManga struct {
	Title       string
	Description string
	Author      string
	CoverImage  string
	Genres      []string
	Language    string
	Status      string
	SourceURL   string
	Chapters    []ScrapedChapter `json:"chapters,omitempty"` 
}
type ScrapedChapter struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Release   string `json:"release,omitempty"` // optional date string
	Number    int    `json:"number,omitempty"`  // optional chapter number
}

type NormalizedChapter struct {
	ID        string `json:"id"`        // unique hash
	Title     string `json:"title"`
	URL       string `json:"url"`
	Release   string `json:"release,omitempty"`
	Number    int    `json:"number,omitempty"`
}
