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

// Insert data into the database
func Insert(data Data) {
	_, err := db.DB.Exec(
		fmt.Sprintf(`
		INSERT INTO %v (
			followers,
			email,
			username,
			repos,
			contributions,
			hireable,
			location,
			organizations,
			website,
			company) VALUES (
				$1::int, $2::varchar, $3::varchar,
				$4::int, $5::int, $6::bool,
				$7::varchar, $8::int, $9::varchar,
				$10::varchar
			);
		`, TableName),
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
	)
	lumber.Error(err, "Failed to update", TableName, "with latest information")
	lumber.Success("Initialized", TableName, "with latest data")
}

// Update the account information in the database
func Update(data Data) {
	_, err := db.DB.Exec(fmt.Sprintf(
		`
		UPDATE %v

		SET followers = $1::int,
		    email = $2::varchar,
			username = $3::varchar,
			repos = $4::int,
			contributions = $5::int,
			hireable = $6::bool,
			location = $7::varchar,
			organizations = $8::int,
			website = $9::varchar,
			company = $10::varchar;
		`, TableName),
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
	)
	lumber.Error(err, "Failed to update", TableName, "with most recent values")
	lumber.Success("Updated", TableName, "with most recent values")
}
