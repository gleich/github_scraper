package main

import (
	"os"
	"time"

	"github.com/gleich/github_scraper/pkg/account"
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/github_scraper/pkg/gh_api"
	"github.com/gleich/github_scraper/pkg/projects"
	"github.com/gleich/lumber"
)

var (
	// All SQL tables used by this microservice
	tables = []db.Table{
		{
			Name:        account.TableName,
			CreateQuery: account.CreateQuery,
		},
		{
			Name:        projects.TableName,
			CreateQuery: projects.CreateQuery,
		},
	}

	// GitHub projects to store in the database
	gitHubProjects = map[string][]string{
		"gleich": {"api", "github_scraper", "fgh", "dots", "lumber", "site-v2", "blackbird"},
		"hackclub":    {"awesome_hackclub_auto"},
	}
)

func main() {
	lumber.ErrNilCheck = false

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
	for owner, repos := range gitHubProjects {
		for _, repo := range repos {
			var (
				rawProjectData = projects.GetData(owner, repo)
				projectData    = projects.CleanData(rawProjectData)
			)
			projects.Insert(projectData)
		}
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
		for owner, repos := range gitHubProjects {
			for _, repo := range repos {
				var (
					rawProjectData = projects.GetData(owner, repo)
					projectData    = projects.CleanData(rawProjectData)
				)
				projects.Update(projectData)
			}
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
