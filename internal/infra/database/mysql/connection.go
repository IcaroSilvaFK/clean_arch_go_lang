package database_provider

import (
	"database/sql"
	"log"
)

func GetConnection(provider, dsn string) *sql.DB {
	db, err := sql.Open(provider, dsn)
	if err != nil {
		log.Fatal(err)
	}

	initializeDatabase(db)

	return db
}

func initializeDatabase(db *sql.DB) {

	_, err := db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		panic(err)
	}
}
