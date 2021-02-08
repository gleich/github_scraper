package projects

import (
	"context"

	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
	"github.com/Matt-Gleich/lumber"
)

func GetData() ApiData {
	lumber.Info("Getting project data")

	var data ApiData
	err := gh_api.Client.Query(context.Background(), &data, nil)
	lumber.Error(err, "Failed to get project data")

	return data
}
