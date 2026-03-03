package core

import (
	"encoding/json"
	"fmt"

	"blackedge-backend/queue"
)

type registryIngestionMessage struct {
	Source string          `json:"source"`
	Data   json.RawMessage `json:"data"`
}

func RunRegistryIngestion() {
	sources := GetAllSources()

	for _, src := range sources {
		fmt.Println("📡 Dispatching source:", src.Name())
		payload, err := src.Fetch()
		if err != nil {
			fmt.Println("❌ Fetch error:", err)
			continue
		}

		queue.Publish("ingestion.raw", registryIngestionMessage{
			Source: src.Name(),
			Data:   json.RawMessage(payload),
		})
	}
}
