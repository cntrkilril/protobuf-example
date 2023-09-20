package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// gomigrate migration resolver
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	// db driver
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	Config struct {
		Postgres
	}

	Postgres struct {
		ConnString      string        `koanf:"connString" validate:"required"`
		MaxOpenConns    int           `koanf:"maxOpenConns" validate:"required"`
		ConnMaxLifetime time.Duration `koanf:"connMaxLifetime" validate:"required"`
		MaxIdleConns    int           `koanf:"maxIdleConns" validate:"required"`
		ConnMaxIdleTime time.Duration `koanf:"connMaxIdleTime" validate:"required"`
		AutoMigrate     bool          `koanf:"autoMigrate"`
		MigrationsPath  string        `koanf:"migrationsPath" validate:"required"`
		DBName          string        `koanf:"dbname" validate:"required"`
	}
)

func InitPsqlDB(cfg Config, l *zap.SugaredLogger) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		l.Error(err)
		return nil
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime * time.Second)

	err = db.Ping()
	if err != nil {
		l.Error(err)
		return nil
	}

	if cfg.Postgres.AutoMigrate {
		migrationDriver, err := postgres.WithInstance(db.DB, &postgres.Config{})
		if err != nil {
			l.Error(err)
			return nil
		}

		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", cfg.Postgres.MigrationsPath),
			cfg.Postgres.DBName,
			migrationDriver,
		)
		if err != nil {
			l.Error(err)
			return nil
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			l.Error(err)
			return nil
		}
	}

	return db
}

func DeferClose(db *sqlx.DB, l *zap.SugaredLogger) {
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			l.Error(err)
			return
		}
	}(db)
}
