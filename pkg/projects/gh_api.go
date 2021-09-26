package projects

import (
	"context"

	"github.com/gleich/github_scraper/pkg/account"
	"github.com/gleich/github_scraper/pkg/gh_api"
	"github.com/gleich/lumber/v2"
	"github.com/shurcooL/githubv4"
)

// Get the data for all my repos
func GetProjectsData(account account.Data) []RepositoryData {
	lumber.Info("Getting repo for my personal projects")

	repos := []RepositoryData{}
	variables := map[string]interface{}{"cursor": (*githubv4.String)(nil)}
	for {
		var data struct {
			Viewer struct {
				Repositories struct {
					Nodes    []RepositoryData
					PageInfo struct {
						HasNextPage bool
						EndCursor   githubv4.String
					}
				} `graphql:"repositories(first: 100, after: $cursor, privacy: PUBLIC, affiliations: OWNER, isFork: false, orderBy: {field: UPDATED_AT, direction: DESC})"`
			}
		}
		err := gh_api.Client.Query(
			context.Background(),
			&data,
			variables,
		)
		if err != nil {
			lumber.Error(err, "Failed to get list of repos")
		}
		variables["cursor"] = githubv4.NewString(data.Viewer.Repositories.PageInfo.EndCursor)

		for _, repo := range data.Viewer.Repositories.Nodes {
			if repo.Owner.Login == account.Username {
				repos = append(repos, repo)
			}
		}

		if !data.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
	}

	lumber.Success("Got data for my personal projects")
	return repos
}

// Get the data for a specific repo
func GetProjectData(owner string, name string) RepositoryData {
	lumber.Info("Getting project data")

	var data struct {
		Repository RepositoryData `graphql:"repository(name: $name, owner: $owner)"`
	}
	err := gh_api.Client.Query(
		context.Background(),
		&data,
		map[string]interface{}{"owner": githubv4.String(owner), "name": githubv4.String(name)},
	)
	if err != nil {
		lumber.Error(err, "Failed to get project data")
	}

	lumber.Success("Got project data")
	return data.Repository
}
