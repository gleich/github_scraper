package projects

import (
	"time"

	"github.com/shurcooL/githubv4"
)

// Data from github
type ApiData struct {
	Viewer struct {
		Repositories struct {
			Nodes []struct {
				NameWithOwner githubv4.String
				Name          githubv4.String
				Owner         struct {
					Login githubv4.String
				}
				URL             githubv4.String
				PrimaryLanguage struct {
					Name  githubv4.String
					Color githubv4.String
				}
				Description    githubv4.String
				PushedAt       githubv4.DateTime
				CreatedAt      githubv4.DateTime
				StargazerCount githubv4.Int
			}
		} `graphql:"repositories(first: 50, privacy: PUBLIC, orderBy: {field: STARGAZERS, direction: DESC}, ownerAffiliations: OWNER)"`
	}
}

// A GitHub Project
type Project struct {
	NameWithOwner string
	Name          string
	Owner         string
	URL           string
	MainLangName  string
	MainLangColor string
	Description   string
	LastPushAt    time.Time
	CreatedAt     time.Time
	Stars         int
}
