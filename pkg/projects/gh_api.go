package projects

import (
	"context"

	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
	"github.com/Matt-Gleich/lumber"
	"github.com/shurcooL/githubv4"
)

func GetData(name string, owner string) ApiData {
	lumber.Info("Getting project info for", owner, "/", name)

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	var data ApiData
	err := gh_api.Client.Query(context.Background(), &data, variables)
	lumber.Error(err, "Failed to get project data for", owner, "/", name)

	return data
}
