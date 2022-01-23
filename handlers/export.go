package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type export struct {
	Transactions []map[string]interface{} `json:"transactions"`
	Schedules    []map[string]interface{} `json:"schedules"`
}

func startExport(uid string, target string) {
	fmt.Printf("Writing documents for user %v into %v.\n", uid, target)
	println()

	var data export

	data.Transactions = exportCollection("transactions", uid)
	data.Schedules = exportCollection("schedules", uid)

	writeFile(target, data)
}

func exportCollection(collection string, uid string) []map[string]interface{} {
	fmt.Printf("Exporting %v...\n", collection)

	documents, _ := FirestoreClient.Collection(collection).Where("uid", "==", uid).Documents(Context).GetAll()
	var results []map[string]interface{}

	for _, document := range documents {
		data := document.Data()
		delete(data, "uid")
		results = append(results, data)
	}

	fmt.Printf("Exported %v %v.\n", len(documents), collection)
	println()

	return results
}

func writeFile(filename string, data export) {
	println("Writing export to file...")

	content, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}

	println("Wrote export to file.")
}
