package db

import (
	"database/sql"
	_ "github.com/lib/pq"
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
	if err := initDBStructure(db); err != nil {
		return err
	}
	DataBase.db = db
	return nil
}

func (DataBase *DataBase) Close() {
	DataBase.db.Close()
}
