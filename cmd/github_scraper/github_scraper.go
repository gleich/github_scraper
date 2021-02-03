package main

import (
	"github.com/Matt-Gleich/github_scraper/pkg/account"
	"github.com/Matt-Gleich/github_scraper/pkg/db"
)

func main() {
	// Connecting to db and resetting tables
	db.Connect()
	db.ResetTable(account.CreateQuery, account.TableName)
}
