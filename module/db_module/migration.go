package db_module

import (
	"context"
	"github.com/samber/lo"
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
	lo.Must0(migrator.Init(ctx))
	return migrator.Migrate(ctx)
}
