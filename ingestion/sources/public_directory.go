package sources

import (
	"blackedge-backend/core"
	"encoding/json"
)

type PublicDirectorySource struct{}

func (p PublicDirectorySource) Name() string {
	return "public_directory"
}

func (p PublicDirectorySource) Fetch() ([]byte, error) {
	data := []map[string]string{
		{"title": "Test Manga", "url": "http://test.com"},
	}
	return json.Marshal(data)
}

func init() {
	core.RegisterSource(PublicDirectorySource{})
}