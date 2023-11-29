package db

import (
	"database/sql"
)

func initDBStructure(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Users(" +
		"discord_id varchar," +
		"PRIMARY KEY(discord_id));"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Wallet(" +
		"wallet_id serial, " +
		"discord_id varchar, " +
		"value bigint, " +
		"PRIMARY KEY(wallet_id), " +
		"FOREIGN KEY(discord_id) REFERENCES Users(discord_id));"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS transactions(" +
		"transaction_id serial, " +
		"from_user varchar, " +
		"to_user varchar, " +
		"value bigint, " +
		"PRIMARY KEY(transaction_id), " +
		"FOREIGN KEY(from_user) REFERENCES Users(discord_id), " +
		"FOREIGN KEY(to_user) REFERENCES Users(discord_id));"); err != nil {
		return err
	}
	return nil
}
