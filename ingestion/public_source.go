package ingestion

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PublicSource struct{}

func (p PublicSource) Name() string {
	return "PublicDirectorySource"
}

func (p PublicSource) BaseURL() string {
	return "https://www.gutenberg.org/browse/categories/1" // Public domain books
}

func (p PublicSource) GetMangaList() ([]ScrapedManga, error) {
	resp, err := http.Get(p.BaseURL())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []ScrapedManga

	doc.Find("li a").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")

		if title != "" && link != "" {
			results = append(results, ScrapedManga{
				Title:     title,
				SourceURL: "https://www.gutenberg.org" + link,
			})
		}
	})

	return results, nil
}
