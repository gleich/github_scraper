package main

import (
	"os"
	"time"

	"github.com/gleich/github_scraper/pkg/account"
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/github_scraper/pkg/gh_api"
	"github.com/gleich/github_scraper/pkg/projects"
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
	extraGitHubProjects = map[string][]string{
		"hackclub": {"awesome_hackclub_auto"},
	}
)

func main() {
	db.Connect()
	gh_api.GenClient()
	resetTables()
	setInitialValues()
	pause()
	// runCycles()
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

	// My projects
	personalRepos := projects.GetProjectsData(githubAccount)
	for _, project := range personalRepos {
		projects.Insert(project)
	}

	// Additional projects
	for owner, repos := range extraGitHubProjects {
		for _, repo := range repos {
			projectData := projects.GetProjectData(owner, repo)
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
		// My projects
		personalRepos := projects.GetProjectsData(githubAccount)
		for _, project := range personalRepos {
			projects.Update(project)
		}

		// Additional projects
		for owner, repos := range extraGitHubProjects {
			for _, repo := range repos {
				projectData := projects.GetProjectData(owner, repo)
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
