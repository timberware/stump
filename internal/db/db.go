package db

import (
	"context"
	"database/sql"
	_ "embed"
	_ "modernc.org/sqlite"

	gen "stump/internal/db/generated"
	"stump/internal/logger"
)

const NAME = "stump.sqlite"

//go:embed schema.sql
var ddl string

func Init() (*sql.DB, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", NAME)
	if err != nil {
		logger.Error("failed to open database: %v", err)
		return nil, err
	}
	logger.Info("%s opened", NAME)

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		logger.Error("failed to create table: %v", err)
		return nil, err
	}
	logger.Info("tables created")

	return db, nil
}

func Queries(db *sql.DB) *gen.Queries {
	return gen.New(db)
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Error("failed to close db: %v", err)
		return
	}
}
