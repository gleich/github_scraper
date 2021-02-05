package projects

// Organize and clean the data from the API
func CleanData(apiData ApiData) Data {
	return Data{
		Name:          string(apiData.Repository.Name),
		URL:           string(apiData.Repository.Url),
		MainLangName:  string(apiData.Repository.PrimaryLanguage.Name),
		MainLangColor: string(apiData.Repository.PrimaryLanguage.Color),
		Description:   string(apiData.Repository.Description),
		Stars:         int(apiData.Repository.Stargazers.TotalCount),
	}
}
