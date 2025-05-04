package db

import (
	"context"
	"errtrack/internal/entity"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func migrateEventTable(migrations *migrate.Migrations) {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			db.NewCreateTable().Model(&entity.Event{})
			return nil
		},
		func(ctx context.Context, db *bun.DB) error {
			db.NewDropTable().Model(&entity.Event{})
			return nil
		},
	)
}
