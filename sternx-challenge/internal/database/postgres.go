package database

import (
	"database/sql"
	"sternx-challenge/config"

	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type postgresDB struct {
	db *sql.DB
}

func NewPostgresDB(cfg *config.Config) (DB, error) {
	dsn := cfg.DB.DataSourceName
	db, err := sql.Open(cfg.DB.DriverName, dsn)
	if err != nil {
		return nil, err
	}
	return &postgresDB{db: db}, nil
}

func (p *postgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args...)
}
