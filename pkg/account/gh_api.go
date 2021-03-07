package account

import (
	"context"

	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
	"github.com/Matt-Gleich/lumber"
)

// Get account data from GitHub
func GetData() ApiData {
	lumber.Info("Getting account data")

	var data ApiData
	err := gh_api.Client.Query(context.Background(), &data, nil)
	if err != nil {
		lumber.Error(err, "Failed to get github account data")
	}

	return data
}
