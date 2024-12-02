package errors_helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/errorreporting"
)

var errorClient *errorreporting.Client

func ReportError(ctx context.Context, clientError error) {
	fmt.Printf("Reporting error: %v", clientError)
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	var err error

	errorClient, err = errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName:    "go-demo-api",
		ServiceVersion: "0.0.0",
		OnError: func(err error) {
			log.Printf("Could not report the error: %v", err)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer errorClient.Close()

	errorClient.Report(errorreporting.Entry{
		Error: clientError,
	})

	log.Print(clientError)
}
