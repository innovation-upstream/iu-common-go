package repo

import (
	"context"
	"database/sql"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gitlab.innovationup.stream/innovation-upstream/iu-common-go/helper"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	log "unknwon.dev/clog/v2"
)

func NewFirestoreClient(ctx context.Context) *firestore.Client {
	projectID := helper.GetProjectID()
	opt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: 5 * time.Minute,
	}))
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatal("%v", errors.WithStack(err))
	}

	return client
}

func NewMySQLDBConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("%+v\n", errors.WithStack(err))
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("%+v\n", errors.WithStack(err))
	}

	return db
}
