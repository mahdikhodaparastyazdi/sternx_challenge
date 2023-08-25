package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type postgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (DB, error) {
	dsn := "user=postgres password=1234 host=localhost dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &postgresDB{db: db}, nil
}

func (p *postgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args...)
}
