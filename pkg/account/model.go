package account

import "github.com/shurcooL/githubv4"

// Data from github
type ApiData struct {
	Viewer struct {
		Login     githubv4.String
		Followers struct {
			TotalCount githubv4.Int
		}
		IsHireable   githubv4.Boolean
		Repositories struct {
			TotalCount githubv4.Int
		}
		Email         githubv4.String
		Location      githubv4.String
		Organizations struct {
			TotalCount githubv4.Int
		}
		WebsiteURL              githubv4.String
		Company                 githubv4.String
		ContributionsCollection struct {
			TotalRepositoryContributions                       githubv4.Int
			TotalRepositoriesWithContributedPullRequests       githubv4.Int
			TotalRepositoriesWithContributedPullRequestReviews githubv4.Int
			TotalRepositoriesWithContributedIssues             githubv4.Int
			TotalRepositoriesWithContributedCommits            githubv4.Int
			TotalPullRequestReviewContributions                githubv4.Int
			TotalCommitContributions                           githubv4.Int
			TotalPullRequestContributions                      githubv4.Int
			TotalIssueContributions                            githubv4.Int
		}
	}
}

// Cleaned data
type Data struct {
	Followers     int
	Email         string
	Username      string
	Repos         int
	Contributions int
	Hireable      bool
	Location      string
	Organizations int
	Website       string
	Company       string
}
