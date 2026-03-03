package workers

import (
	"encoding/json"
	"log"

	"blackedge-backend/queue"
)

type ingestQueueMessage struct {
	Source string          `json:"source"`
	Data   json.RawMessage `json:"data"`
}

// StartIngestWorker subscribes to the ingestion queue and processes every batch.
func StartIngestWorker() {
	queue.Subscribe("ingestion.raw", func(payload []byte) {
		var msg ingestQueueMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println("ingest worker: failed to parse payload:", err)
			return
		}

		log.Println("ingest worker: processing source", msg.Source)
		Process(msg.Data, msg.Source)
	})
}
