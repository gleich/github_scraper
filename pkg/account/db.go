package account

import (
	"fmt"

	"github.com/Matt-Gleich/github_scraper/pkg/db"
	"github.com/Matt-Gleich/lumber"
)

// Table name for the account info
const TableName = "github_account"

// Create query
var CreateQuery = fmt.Sprintf(
	`
CREATE TABLE %v (
	followers int,
	email varchar(255),
	username varchar(255),
	repos int,
	contributions int,
	hireable bool,
	location varchar(255),
	organizations int,
	website varchar(255),
	company varchar(255)
);`, TableName,
)

func Insert(data Data) {
	// Inserting data
	_, err := db.DB.Exec(fmt.Sprintf(
		` INSERT INTO %v (followers,
			email,
			username,
			repos,
			contributions,
			hireable,
			location,
			organizations,
			website,
			company) VALUES (
				%v, '%s', '%s', %v, %v, %v, '%s', %v, '%s', '%s'
			);
		`,
		TableName,
		data.Followers,
		data.Email,
		data.Username,
		data.Repos,
		data.Contributions,
		data.Hireable,
		data.Location,
		data.Organizations,
		data.Website,
		data.Company,
	))
	lumber.Error(err, "Failed to update", TableName, "with latest information")
	lumber.Success("Updated", TableName, "with latest data")
}
