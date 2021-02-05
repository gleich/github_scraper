package main

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/github_scraper/pkg/account"
	"github.com/Matt-Gleich/github_scraper/pkg/db"
	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
	"github.com/Matt-Gleich/github_scraper/pkg/projects"
)

// All SQL tables used by this microservice
var tables = []db.Table{
	{
		Name:        account.TableName,
		CreateQuery: account.CreateQuery,
	},
	{
		Name:        projects.TableName,
		CreateQuery: projects.CreateQuery,
	},
}

// All projects
var gitHubProjects = []projects.Project{
	{Name: "fgh", Owner: "Matt-Gleich"},
}

func main() {
	// Connecting to db and resetting tables
	db.Connect()
	for _, table := range tables {
		db.HardResetTable(table.CreateQuery, table.Name)
	}

	gh_api.GenClient()

	for {
		// Getting account information
		rawAccountData := account.GetData()
		formattedAccountData := account.CleanData(rawAccountData)
		db.ResetTable(account.TableName)
		account.Insert(formattedAccountData)

		for _, project := range gitHubProjects {
			rawProjectData := projects.GetData(project.Name, project.Owner)
			fmt.Println(rawProjectData)
		}

		time.Sleep(5 * time.Minute)
	}
}
