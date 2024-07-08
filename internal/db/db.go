package db

import (
	"context"
	"database/sql"
	_ "embed"
	_ "modernc.org/sqlite"
	"path/filepath"
	gen "stump/internal/db/generated"
	"stump/internal/logger"
	"stump/internal/utils"
)

const dbName = "stump.sqlite"

//go:embed schema.sql
var ddl string

func Init() (*sql.DB, error) {
	ctx := context.Background()

	dbPath, err := utils.GetAppPath()
	if err != nil {
		logger.Error("failed to get database path: %v", err)
		return nil, err
	}
	dbPath = filepath.Join(dbPath, dbName)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info("Database ready")

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		logger.Error("failed to create table: %v", err)
		return nil, err
	}

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
