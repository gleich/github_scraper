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

func main() {
	db.Connect()
	gh_api.GenClient()
	resetTables()
	setInitialValues()
	pause()
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
	var (
		rawAccountData = account.GetData()
		githubAccount  = account.CleanData(rawAccountData)
	)
	account.Insert(githubAccount)

	// Projects
	var (
		projectData    = projects.GetData()
		githubProjects = projects.CleanData(projectData)
	)
	for _, project := range githubProjects {
		projects.Insert(project)
	}
}

// Run the update cycles
func runCycles() {
	for {
		// Account information
		var (
			rawAccountData = account.GetData()
			githubAccount  = account.CleanData(rawAccountData)
		)
		account.Update(githubAccount)

		// Projects
		var (
			projectData    = projects.GetData()
			githubProjects = projects.CleanData(projectData)
		)
		for _, project := range githubProjects {
			projects.Update(project)
		}

		pause()
	}
}

// Provide a time buffer
func pause() {
	// Pausing for next run
	if os.Getenv("DEV_UPDATE_TIME") == "true" {
		time.Sleep(15 * time.Second)
	} else {
		time.Sleep(5 * time.Minute)
	}
}
