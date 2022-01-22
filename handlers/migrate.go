package handlers

import (
	"fmt"

	"cloud.google.com/go/firestore"
)

func startMigration(from string, to string, dryRun bool) {
	println("Running migration with the following parameters:\n")
	println("From:", from)
	println("To:", to)

	if dryRun {
		println("\n[!] Dry run is active. Won't send any writing requests.")
	}

	println()
	migrateCollection("transactions", from, to, dryRun)
	println()
	migrateCollection("schedules", from, to, dryRun)
}

func migrateCollection(collection string, from string, to string, dryRun bool) {
	fmt.Printf("Migrating %v...\n", collection)

	schedules, _ := Client.Collection(collection).Where("uid", "==", from).Documents(Context).GetAll()
	scheduleCount := len(schedules)

	for i := 0; i < scheduleCount; i++ {
		doc := schedules[i]

		data := doc.Data()
		data["uid"] = to

		updateDocument(doc.Ref, data, dryRun)
	}

	fmt.Printf("Migrated %v %v.\n", scheduleCount, collection)
}

func updateDocument(documentRef *firestore.DocumentRef, data map[string]interface{}, dryRun bool) {
	if dryRun {
		return
	}

	documentRef.Set(Context, data)
}
