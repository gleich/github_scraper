package projects

import (
	"context"

	"github.com/gleich/github_scraper/pkg/gh_api"
	"github.com/gleich/lumber/v2"
	"github.com/shurcooL/githubv4"
)

// Get the data for a project
func GetData(owner string, name string) ApiData {
	lumber.Info("Getting project data")

	var data ApiData
	err := gh_api.Client.Query(
		context.Background(),
		&data,
		map[string]interface{}{"owner": githubv4.String(owner), "name": githubv4.String(name)},
	)
	if err != nil {
		lumber.Error(err, "Failed to get project data")
	}

	return data
}
