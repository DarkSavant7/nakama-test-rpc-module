package database

import (
	"context"
	"dark.savant.org/nakama-test-rpc-module/structures"
	"database/sql"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
)

func SaveToDB(ctx context.Context, db *sql.DB, logger runtime.Logger, request structures.Request) error {
	logger.Debug("[TEST RPC MODULE]: Inserting new log")
	result, err := db.ExecContext(ctx, INSERT_LOG, request.Type, request.Version, request.Hash)

	if err != nil {
		logger.Error("[TEST RPC MODULE]: Error occurred while inserting new log: %s", err.Error())
		return err
	}

	insertedId, _ := result.LastInsertId()
	logger.Debug("[TEST RPC MODULE]: Log inserted successfully. Id: %v", insertedId)

	return nil
}

func InitDB(ctx context.Context, db *sql.DB, logger runtime.Logger) {
	logger.Debug("[TEST RPC MODULE]: Initializing DB for the TEST RPC module")
	_, err := db.ExecContext(ctx, CREATE_TABLE)

	if err != nil {
		errorMessage := fmt.Sprintf("[TEST RPC MODULE]: Error occurred while initializing log table for the RPC test module: %s\nAborting", err.Error())
		logger.Error(errorMessage)
		panic(errorMessage)
	}

	logger.Debug("[TEST RPC MODULE]: DB for the TEST RPC module initialized successfully")
}
