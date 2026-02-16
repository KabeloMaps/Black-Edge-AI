package ingestion

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type MangaSource struct {
	Name string
	URL string
}

func ScrapeMangaList(source MangaSource) ([]string, error) {

	resp, err := http.Get(source.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var mangaLinks []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			mangaLinks = append(mangaLinks, link)
		}
	})

	return mangaLinks, nil
}