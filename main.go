// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

//Task
type Task struct {
	Description string
}

func main() {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "CHANGE_ME"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "Task"
	// Sets the name/ID for the new entity.
	name := "sampletask1"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Buy milk",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

	gettask := Task{
		Description: "Buy milk",
	}

	if err := client.Get(ctx, taskKey, &gettask); err != nil {
		log.Fatalf("Failed to get task: %v", err)
	}

	fmt.Printf("Get %v: %v\n", taskKey, task.Description)
}
