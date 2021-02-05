package projects

import "github.com/shurcooL/githubv4"

// "dumb" project data
type Project struct {
	Name  string
	Owner string
}

// Data from github
type ApiData struct {
	Repository struct {
		Name       githubv4.String
		Url        githubv4.String
		Stargazers struct {
			TotalCount githubv4.Int
		}
		Description     githubv4.String
		PrimaryLanguage struct {
			Color githubv4.String
			Name  githubv4.String
		}
	} `graphql:"repository(owner: $owner, name: $name)"`
}

// A GitHub Project
type Data struct {
	Name          string
	URL           string
	MainLangName  string
	MainLangColor string
	Description   string
	Stars         int
}
