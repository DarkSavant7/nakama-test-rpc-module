package main

import (
	"context"
	"dark.savant.org/nakama-test-rpc-module/database"
	"dark.savant.org/nakama-test-rpc-module/rpc"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("[TEST RPC MODULE]: initializing....")
	database.InitDB(ctx, db, logger)

	//register our RPC function
	if err := initializer.RegisterRpc("test_rpc_func_id", rpc.TestRpcFunction); err != nil {
		logger.Error("[TEST RPC MODULE]: Unable to register: %v", err)
		return err
	}

	return nil
}
