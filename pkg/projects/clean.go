package projects

// Organize and clean the data from the API
func CleanData(apiData ApiData) []Project {
	projects := []Project{}
	for _, project := range apiData.Viewer.Repositories.Nodes {
		projects = append(projects, Project{
			NameWithOwner: string(project.NameWithOwner),
			Name:          string(project.Name),
			Owner:         string(project.Owner.Login),
			URL:           string(project.URL),
			MainLangName:  string(project.PrimaryLanguage.Name),
			MainLangColor: string(project.PrimaryLanguage.Color),
			Description:   string(project.Description),
			LastPushAt:    project.PushedAt.Time,
			CreatedAt:     project.CreatedAt.Time,
			Stars:         int(project.StargazerCount),
		})
	}
	return projects
}
