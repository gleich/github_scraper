package projects

import (
	"github.com/shurcooL/githubv4"
)

// Raw repository data from github
type RepositoryData struct {
	NameWithOwner string
	Name          string
	Owner         struct {
		Login string
	}
	URL             string
	PrimaryLanguage struct {
		Name  string
		Color string
	}
	Description    string
	PushedAt       githubv4.DateTime
	CreatedAt      githubv4.DateTime
	StargazerCount githubv4.Int
}
