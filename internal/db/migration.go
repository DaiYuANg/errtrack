package db

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/fx"
)

var migrationModule = fx.Module("migration",
	fx.Provide(newMigration),
	fx.Invoke(migrateEventTable, runMigrate),
)

func newMigration() *migrate.Migrations {
	return migrate.NewMigrations()
}

func runMigrate(db *bun.DB, migrations *migrate.Migrations) (*migrate.MigrationGroup, error) {
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, migrations)
	migrator.Init(ctx)
	return migrator.Migrate(context.Background())
}
