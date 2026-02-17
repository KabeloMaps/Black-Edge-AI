package ingestion

import "fmt"

func RunIngestion(source SourceAdapter) ([]NormalizedManga, error) {
	fmt.Println("🔁 Ingesting from:", source.Name())

	rawData, err := source.GetMangaList()
	if err != nil {
		return nil, err
	}

	fmt.Println("📥 Raw items scraped:", len(rawData))

	normalized := NormalizeManga(rawData, source.Name())
	fmt.Println("📦 Normalized items:", len(normalized))

	return normalized, nil
}
