package main

import (
	"github.com/Matt-Gleich/github_scraper/pkg/account"
	"github.com/Matt-Gleich/github_scraper/pkg/db"
	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
)

func main() {
	// Connecting to db and resetting tables
	db.Connect()
	db.ResetTable(account.CreateQuery, account.TableName)

	client := gh_api.GenClient()

	// Getting account information
	account.GetData(client)
}
