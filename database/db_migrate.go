package database

import "database/sql"

func runMigrations(db *sql.DB) error {
	err := verifyCrawlerTable(db)
	if err != nil {
		return err
	}

	return err
}

func verifyCrawlerTable(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS crawler")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS papers (
    		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    		name varchar(200) NOT NULL,
   			company varchar(200) NOT NULL,
				marketValue FLOAT NOT NULL,
    		dailyVariation VARCHAR(15),
    		url VARCHAR(400) NOT NULL
		);`)
	return err
}
