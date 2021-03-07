package projects

import (
	"context"

	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
	"github.com/Matt-Gleich/lumber"
	"github.com/shurcooL/githubv4"
)

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
