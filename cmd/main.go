package main

import (
	"blackedge-backend/core"
	"blackedge-backend/ingestion/workers"
	"blackedge-backend/queue"
	"blackedge-backend/storage"
	"log"
)

func main() {

	storage.ConnectMongo("mongodb://127.0.0.1:27017", "blackedge")
	queue.ConnectRedis()

	workers.StartIngestWorker()

	go core.RunRegistryIngestion()

	log.Println("🚀 BlackEdge Core Online")

	select {} // keep alive
}