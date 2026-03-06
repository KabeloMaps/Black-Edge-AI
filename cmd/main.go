package main

import (
	"blackedge-backend/core"
	"blackedge-backend/ingestion/workers"
	"blackedge-backend/queue"
	"blackedge-backend/registry"
	"blackedge-backend/storage"
	"fmt"
	"log"
)

func main() {

	fmt.Println("🔍 Starting BlackEdge Core...")
	instruments := registry.GetAllInstruments()
	fmt.Println("📊 BlackEdge Instrument Registry")

	for _, inst := range instruments {
		fmt.Println(inst.Symbol)
	}

	storage.ConnectMongo("mongodb://127.0.0.1:27017", "blackedge")
	queue.ConnectRedis()

	workers.StartIngestWorker()

	go core.RunRegistryIngestion()

	log.Println("🚀 BlackEdge Core Online")

	for _, i := range instruments{
		fmt.Println(i.Symbol, "-", i.Type)
	}

	//select {} // keep alive
}