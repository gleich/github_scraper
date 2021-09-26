package main

import (
	"os"
	"time"

	"github.com/gleich/github_scraper/pkg/account"
	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/github_scraper/pkg/gh_api"
	"github.com/gleich/github_scraper/pkg/projects"
	"github.com/gleich/lumber/v2"
	"github.com/go-co-op/gocron"
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
	cycle(account.Insert, projects.Insert)
	scheduledCycles()
}

// Reset tables in the database
func resetTables() {
	for _, table := range tables {
		db.HardResetTable(table.CreateQuery, table.Name)
	}
}

func cycle(accountSQLFunc func(account.Data), projectsSQLFunc func(projects.RepositoryData)) {
	// Account information
	var (
		rawAccountData = account.GetData()
		githubAccount  = account.CleanData(rawAccountData)
	)
	accountSQLFunc(githubAccount)

	// My projects
	personalRepos := projects.GetProjectsData(githubAccount)
	for _, project := range personalRepos {
		projectsSQLFunc(project)
	}

	// Additional projects
	for owner, repos := range extraGitHubProjects {
		for _, repo := range repos {
			projectData := projects.GetProjectData(owner, repo)
			projectsSQLFunc(projectData)
		}
	}
}

func scheduledCycles() {
	s := gocron.NewScheduler(time.UTC)
	scheduledCycle := func() {
		cycle(account.Update, projects.Update)
	}

	if os.Getenv("DEV_UPDATE_TIME") == "true" {
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

	lumber.Success("Setup scheduler")
	s.StartBlocking()
}
