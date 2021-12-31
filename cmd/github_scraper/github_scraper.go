package main

import (
	"github.com/gleich/github_scraper/pkg/api"
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/lumber/v2"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		lumber.Fatal(err, "Failed to connect to database")
	}

	err = database.AutoMigrate(&db.Account{})
	if err != nil {
		lumber.Fatal(err, "Failed to auto migrate account")
	}

	client := api.Client()

	init_account, err := api.Account(client)
	if err != nil {
		lumber.Fatal(err, "Failed to fetch initial account information")
	}
	database.Create(&init_account)
}
