package main

import (
	"os"
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
	{Name: "api", Owner: "Matt-Gleich"},
	{Name: "cihat", Owner: "Matt-Gleich"},
	{Name: "awesome_hackclub_auto", Owner: "hackclub"},
	{Name: "dots", Owner: "Matt-Gleich"},
	{Name: "nuke", Owner: "Matt-Gleich"},
	{Name: "import_sorter", Owner: "fluttercommunity"},
	{Name: "lumber", Owner: "Matt-Gleich"},
}

func main() {
	// Connecting to db and creating tables if they
	// don't already exist
	db.Connect()
	for _, table := range tables {
		db.SafeCreate(table.CreateQuery, table.Name)
	}

	gh_api.GenClient()

	for {
		// Getting account information
		rawAccountData := account.GetData()
		formattedAccountData := account.CleanData(rawAccountData)
		db.ResetTable(account.TableName)
		account.Insert(formattedAccountData)

		// Getting project information
		db.ResetTable(projects.TableName)
		for _, project := range gitHubProjects {
			rawProjectData := projects.GetData(project.Name, project.Owner)
			formattedProjectData := projects.CleanData(rawProjectData)
			projects.Insert(formattedProjectData)
		}

		if os.Getenv("DEV_UPDATE_TIME") == "true" {
			time.Sleep(15 * time.Second)
		} else {
			time.Sleep(5 * time.Minute)
		}
	}
}
