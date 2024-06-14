package dbutils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/ethereum/go-ethereum/log"
	_ "github.com/lib/pq"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
)

func Open(urlDatabase string) *db.Client {
	sqldb, err := sql.Open("postgres", urlDatabase)
	if err != nil {
		panic(fmt.Errorf("Could not connect to db, error: %w", err))
	}

	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Minute * 30)
	drv := entsql.OpenDB("postgres", sqldb)

	return db.NewClient(db.Driver(drv))
}

func MigrateDB(client *db.Client, ctx context.Context) {
	if err := client.Schema.Create(ctx); err != nil {
		panic(fmt.Errorf("failed creating schema resources: %w", err))
	}
	log.Info("Migrated DB successfully")
}
