package db

import (
	"database/sql"
	"github.com/alexlast/bunzap"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

var Module = fx.Module("db", fx.Provide(createDBConnection), migrationModule)

func createDBConnection(logger *zap.Logger) *bun.DB {
	dsn := "postgres://root:root@localhost:5432/errtrack?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	queryHook := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	db.AddQueryHook(queryHook)
	db.AddQueryHook(
		bunzap.NewQueryHook(bunzap.QueryHookOptions{
			Logger:       logger,
			SlowDuration: 200 * time.Millisecond,
		}),
	)
	return db
}
