package repo

import (
	"context"
	"database/sql"
	"os"

	"cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
	log "unknwon.dev/clog/v2"
)

func NewFirestoreClient(ctx context.Context) *firestore.Client {
	projectID := os.Getenv("GCP_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal("%v", err)
	}

	return client
}

func NewMySQLDBConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("%+v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("%+v\n", err)
	}

	return db
}
