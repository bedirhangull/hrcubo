package postgres

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// migrationsFS is a filesystem that embeds the migrations folder
//
//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	poolConfig.ConnConfig.ConnectTimeout = 5 * time.Second

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %v", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &DB{
		Pool:         db,
		QueryBuilder: &psql,
		url:          url,
	}, nil
}

func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d *DB) ErrorCode(err error) string {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return "UNKNOWN"
	}
	return pgErr.Code
}

func (d *DB) Close() {
	d.Pool.Close()
}
