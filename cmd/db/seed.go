package main

import (
	"log"

	"github.com/loongy/jaguar/models"
	"github.com/loongy/jaguar/nulls"
	"github.com/loongy/jaguar/storage/db"
)

func seed(db *db.DB) error {
	log.Println("seeding...")
	_, err := seedUserTable(db)
	if err != nil {
		return err
	}
	log.Println("seeding done.")
	return nil
}

func seedUserTable(db *db.DB) (models.Users, error) {
	log.Print("seeding user table... ")

	users := models.Users{
		&models.User{
			EmailAddress: nulls.ValidString("alice@example.com"),
		},
		&models.User{
			EmailAddress: nulls.ValidString("bob@example.com"),
		},
	}

	for i, user := range users {
		user, err := db.InsertUser(user)
		if err != nil {
			log.Println("Insert user error: '", err, "'")
			continue
		}
		users[i] = user
	}

	log.Println("done (", len(users), " seeded).")
	return users, nil
}
