package data

import (
	"database/sql"
	"os"
)

func OpenDb() (*sql.DB, error) {
	return sql.Open("sqlite3", "./forum.db")
}

func InitTables(db *sql.DB) error {
	sqlData, err := os.ReadFile("./data/shema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlData))
	if err != nil {
		return err
	}

	return nil
}
