package gh_api

import (
	"context"
	"os"

	"github.com/Matt-Gleich/lumber"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Generate a github api v4 client
func GenClient() *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_PAT_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	lumber.Info("Created githubv4 client")
	return githubv4.NewClient(httpClient)
}
