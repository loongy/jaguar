package main

import (
	"log"

	"github.com/loongy/jaguar-template/storage/db"
)

func setup(db *db.DB) error {
	log.Println("setting up...")
	if err := setupUserTable(db); err != nil {
		return err
	}
	log.Println("setup done.")
	return nil
}

func setupUserTable(db *db.DB) error {
	log.Print("setting up user table... ")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            BIGSERIAL PRIMARY KEY NOT NULL,
			created_at    TIMESTAMP NOT NULL,
			updated_at    TIMESTAMP,
			deleted_at    TIMESTAMP,
			email_address VARCHAR,
		)`)
	if err != nil {
		return err
	}
	log.Println("done.")
	return nil
}
