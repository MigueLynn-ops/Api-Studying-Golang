package persistency

import (
	"Api-Aula_1/config"
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Cfg.FormatDSN())
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Println("Error pinging the database:", err)
		return nil, err
	}

	log.Println("Database connection established successfully")
	return db, nil
}
