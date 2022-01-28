package handlers

import (
	"encoding/json"
	"log"
	"os"
)

func startExport(uid string, target string) {
	log.Printf("Writing documents for user %v into %v.\n", uid, target)
	println()

	var data FileExport

	log.Println("Exporting transactions...")
	data.Transactions = exportTransactions(uid)
	log.Printf("Exported %v transactions.\n", len(data.Transactions))

	println()

	log.Println("Exporting schedules...")
	data.Schedules = exportSchedules(uid)
	log.Printf("Exported %v schedules.\n", len(data.Schedules))

	println()

	log.Println("Writing export to file...")
	writeFile(target, data)
	log.Println("Wrote export to file.")
}

func exportTransactions(uid string) []TransactionExport {
	documents, _ := FirestoreClient.Collection("transactions").Where("uid", "==", uid).Documents(Context).GetAll()
	var results []TransactionExport

	for _, document := range documents {
		var documentExport TransactionExport
		documentExport.fromJSON(document.Data())

		results = append(results, documentExport)
	}

	return results
}

func exportSchedules(uid string) []ScheduleExport {
	documents, _ := FirestoreClient.Collection("schedules").Where("uid", "==", uid).Documents(Context).GetAll()
	var results []ScheduleExport

	for _, document := range documents {
		var documentExport ScheduleExport
		documentExport.fromJSON(document.Data())

		results = append(results, documentExport)
	}

	return results
}

func writeFile(filename string, data FileExport) {
	content, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}
}
