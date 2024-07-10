package config

import (
	"database/sql"

	"github.com/RohanSinghCode/gin-practice/internal/database"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	DB               *database.Queries
	ConnectionString string
}

func (db *DbConnection) ConnectToDb() (*sql.DB, error) {
	conn, err := sql.Open("postgres", db.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.DB = database.New(conn)

	return conn, nil
}
