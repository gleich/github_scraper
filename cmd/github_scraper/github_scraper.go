package main

import (
	"os"
	"time"

	"github.com/gleich/github_scraper/pkg/api"
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/lumber/v2"
	"github.com/go-co-op/gocron"
	"github.com/shurcooL/githubv4"
	"gorm.io/gorm"
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

	accountAPIData, err := api.Account(client)
	if err != nil {
		lumber.Fatal(err, "Failed to fetch initial account information")
	}

	var account db.Account
	database.FirstOrCreate(&account, accountAPIData)

	s := gocron.NewScheduler(time.UTC)

	scheduledCycle := func() { cycle(database, client, &account) }
	if os.Getenv("DEV") == "true" {
		_, err := s.Every(5).Seconds().Do(scheduledCycle)
		if err != nil {
			lumber.Error(err, "Failed to start scheduled cycle for every 5 seconds")
		}
	} else {
		_, err := s.Every(5).Minutes().Do(scheduledCycle)
		if err != nil {
			lumber.Error(err, "Failed to start scheduled cycle for every 5 minutes")
		}
	}

	s.StartBlocking()
}

func cycle(database *gorm.DB, client *githubv4.Client, account *db.Account) {
	new_account_data, err := api.Account(client)
	if err != nil {
		lumber.Error(err, "Failed to load new account data")
	}
	database.Model(&account).Updates(new_account_data)

	lumber.Success("Cycle completed")
}
