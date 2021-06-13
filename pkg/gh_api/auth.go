package gh_api

import (
	"context"
	"os"

	"github.com/gleich/lumber"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// A GitHub api v4 client
var Client *githubv4.Client

// Generate a github api v4 client
func GenClient() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_PAT_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	lumber.Info("Created githubv4 client")
	Client = githubv4.NewClient(httpClient)
}
