package db

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Followers     uint
	Email         string
	Username      string
	Repos         uint
	Contributions uint
	Hireable      bool
	Location      string
	Organizations uint
	Website       string
	Company       string
}
