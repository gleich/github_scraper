package projects

import (
	"fmt"

	"github.com/Matt-Gleich/github_scraper/pkg/db"
	"github.com/Matt-Gleich/lumber"
)

// Table name for projects
const TableName = "github_projects"

// Create query
var CreateQuery = fmt.Sprintf(
	`
CREATE TABLE %v (
	pname_with_owner varchar(255),
	pname varchar(255),
	powner varchar(255),
	url varchar(255),
	main_lang_name varchar(255),
	main_lang_color varchar(255),
	description varchar(255),
	stars int
);`, TableName,
)

// Insert data into the database
func Insert(data Data) {
	_, err := db.DB.Exec(fmt.Sprintf(
		` INSERT INTO %v (
			pname_with_owner,
			pname,
			powner,
			url,
			main_lang_name,
			main_lang_color,
			description,
			stars) VALUES (
				'%s', '%s', '%s', '%s', '%s', '%s', '%s', %v
			);
		`,
		TableName,
		data.NameWithOwner,
		data.Name,
		data.Owner,
		data.URL,
		data.MainLangName,
		data.MainLangColor,
		data.Description,
		data.Stars,
	))
	lumber.Error(err, "Failed to update", TableName, "with latest values for", data.NameWithOwner)
	lumber.Success("Initialized", TableName, "with latest data for", data.NameWithOwner)
}

// Update the values for a certain project
func Update(data Data) {
	_, err := db.DB.Exec(fmt.Sprintf(
		`
		UPDATE %v

		SET pname_with_owner = '%s',
			pname = '%s',
			powner = '%s',
			url = '%s',
			main_lang_name = '%s',
			main_lang_color = '%s',
			description = '%s',
			stars = %v

		WHERE pname_with_owner = '%s';
		`,
		TableName,
		data.NameWithOwner,
		data.Name,
		data.Owner,
		data.URL,
		data.MainLangName,
		data.MainLangColor,
		data.Description,
		data.Stars,
		data.NameWithOwner,
	))
	lumber.Error(err, "Failed to update", TableName, "with most recent values for", data.NameWithOwner)
	lumber.Success("Updated", TableName, "with most recent values for", data.NameWithOwner)
}
