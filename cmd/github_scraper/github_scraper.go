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
	db.Connect()
	gh_api.GenClient()
	resetTables()
	setInitialValues()
	runCycles()
}

// Reset tables in the database
func resetTables() {
	for _, table := range tables {
		db.HardResetTable(table.CreateQuery, table.Name)
	}
}

// Set initial values in the database
func setInitialValues() {
	// Account information
	rawAccountData := account.GetData()
	formattedAccountData := account.CleanData(rawAccountData)
	account.Insert(formattedAccountData)

	// Projects
	for _, project := range gitHubProjects {
		rawProjectData := projects.GetData(project.Name, project.Owner)
		formattedProjectData := projects.CleanData(rawProjectData)
		projects.Insert(formattedProjectData)
	}
}

// Run the update cycles
func runCycles() {
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

		// Pausing for next run
		if os.Getenv("DEV_UPDATE_TIME") == "true" {
			time.Sleep(15 * time.Second)
		} else {
			time.Sleep(5 * time.Minute)
		}
	}
}
