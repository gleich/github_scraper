package main

import (
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/lumber/v2"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		lumber.Fatal(err, "Failed to connect to database")
	}
}
