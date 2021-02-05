package account

// Organize and clean the data from the API
func CleanData(apiData ApiData) Data {
	return Data{
		Username:      string(apiData.Viewer.Login),
		Followers:     int(apiData.Viewer.Followers.TotalCount),
		Hireable:      bool(apiData.Viewer.IsHireable),
		Repos:         int(apiData.Viewer.Repositories.TotalCount),
		Email:         string(apiData.Viewer.Email),
		Location:      string(apiData.Viewer.Location),
		Organizations: int(apiData.Viewer.Organizations.TotalCount),
		Website:       string(apiData.Viewer.WebsiteURL),
		Company:       string(apiData.Viewer.Company),
		Contributions: (int(apiData.Viewer.ContributionsCollection.TotalRepositoryContributions) +
			int(apiData.Viewer.ContributionsCollection.TotalRepositoriesWithContributedPullRequests) +
			int(apiData.Viewer.ContributionsCollection.TotalRepositoriesWithContributedPullRequestReviews) +
			int(apiData.Viewer.ContributionsCollection.TotalRepositoriesWithContributedIssues) +
			int(apiData.Viewer.ContributionsCollection.TotalRepositoriesWithContributedCommits) +
			int(apiData.Viewer.ContributionsCollection.TotalPullRequestReviewContributions) +
			int(apiData.Viewer.ContributionsCollection.TotalCommitContributions) +
			int(apiData.Viewer.ContributionsCollection.TotalPullRequestContributions) +
			int(apiData.Viewer.ContributionsCollection.TotalIssueContributions)),
	}
}
