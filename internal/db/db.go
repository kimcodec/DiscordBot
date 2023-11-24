package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	config *DBConfig
	db     *sql.DB
}

func NewDB(c *DBConfig) *DataBase {
	return &DataBase{
		config: c,
	}
}

func (DataBase *DataBase) Open() error {
	db, err := sql.Open("postgres", DataBase.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	DataBase.db = db
	log.Println("Database connected successfully")
	return nil
}

func (DataBase *DataBase) Close() {
	DataBase.db.Close()
}
