package api

import (
	"context"

	"github.com/gleich/github_scraper/pkg/db"
	"github.com/shurcooL/githubv4"
)

func Account(client *githubv4.Client) (db.Account, error) {
	query := struct {
		Viewer struct {
			Login     string
			Followers struct {
				TotalCount int
			}
			IsHireable   bool
			Repositories struct {
				TotalCount int
			}
			Email         string
			Location      string
			Organizations struct {
				TotalCount int
			}
			WebsiteURL              string
			Company                 string
			ContributionsCollection struct {
				TotalRepositoryContributions                       int
				TotalRepositoriesWithContributedPullRequests       int
				TotalRepositoriesWithContributedPullRequestReviews int
				TotalRepositoriesWithContributedIssues             int
				TotalRepositoriesWithContributedCommits            int
				TotalPullRequestReviewContributions                int
				TotalCommitContributions                           int
				TotalPullRequestContributions                      int
				TotalIssueContributions                            int
			}
		}
	}{}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		return db.Account{}, nil
	}

	return db.Account{
		Followers: uint(query.Viewer.Followers.TotalCount),
		Email:     query.Viewer.Email,
		Username:  query.Viewer.Login,
		Repos:     uint(query.Viewer.Repositories.TotalCount),
		Contributions: uint(
			query.Viewer.ContributionsCollection.TotalRepositoryContributions + query.Viewer.ContributionsCollection.TotalRepositoriesWithContributedPullRequests + query.Viewer.ContributionsCollection.TotalRepositoriesWithContributedPullRequestReviews + query.Viewer.ContributionsCollection.TotalRepositoriesWithContributedIssues + query.Viewer.ContributionsCollection.TotalRepositoriesWithContributedCommits + query.Viewer.ContributionsCollection.TotalPullRequestReviewContributions + query.Viewer.ContributionsCollection.TotalCommitContributions + query.Viewer.ContributionsCollection.TotalPullRequestContributions + query.Viewer.ContributionsCollection.TotalIssueContributions,
		),
		Hireable:      query.Viewer.IsHireable,
		Location:      query.Viewer.Location,
		Organizations: uint(query.Viewer.Organizations.TotalCount),
		Website:       query.Viewer.WebsiteURL,
		Company:       query.Viewer.Company,
	}, nil
}
