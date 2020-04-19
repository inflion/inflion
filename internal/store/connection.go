package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Connection struct {
	db *sql.DB
}

func NewDbConnection() (*sql.DB, error) {
	con := &Connection{}
	con, err := con.Connection()
	if err != nil {
		return nil, err
	}
	return con.db, nil
}

func NewDbtx() (DBTX, error) {
	con := &Connection{}
	con, err := con.Connection()
	if err != nil {
		return nil, err
	}
	return con.db, nil
}

func (c Connection) Connection() (*Connection, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)
	pool.SetConnMaxLifetime(5 * time.Minute)

	c.db = pool

	return &c, nil
}
