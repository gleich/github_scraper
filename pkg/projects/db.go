package projects

import (
	"fmt"

	"github.com/gleich/github_scraper/pkg/db"
	"github.com/gleich/lumber"
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
	last_push_at varchar(255),
	created_at varchar(255),
	stars int
);`, TableName,
)

// Insert data into the database
func Insert(data Project) {
	_, err := db.DB.Exec(fmt.Sprintf(`
		INSERT INTO %v (
			pname_with_owner,
			pname,
			powner,
			url,
			main_lang_name,
			main_lang_color,
			description,
			last_push_at,
			created_at,
			stars) VALUES (
				$1::varchar, $2::varchar, $3::varchar,
				$4::varchar, $5::varchar, $6::varchar,
				$7::varchar, $8::varchar, $9::varchar,
				$10::int
			);
		`, TableName),
		data.NameWithOwner,
		data.Name,
		data.Owner,
		data.URL,
		data.MainLangName,
		data.MainLangColor,
		data.Description,
		data.LastPushAt,
		data.CreatedAt,
		data.Stars,
	)
	if err != nil {
		lumber.Error(
			err,
			"Failed to update",
			TableName,
			"with latest values for",
			data.NameWithOwner,
		)
	}
	lumber.Success("Initialized", TableName, "with latest data for", data.NameWithOwner)
}

// Update the values for a certain project
func Update(data Project) {
	_, err := db.DB.Exec(fmt.Sprintf(`
		UPDATE %v

		SET pname_with_owner = $1::varchar,
			pname = $2::varchar,
			powner = $3::varchar,
			url = $4::varchar,
			main_lang_name = $5::varchar,
			main_lang_color = $6::varchar,
			description = $7::varchar,
			last_push_at = $8::varchar,
			created_at = $9::varchar,
			stars = $10::int

		WHERE pname_with_owner = $11::varchar;
		`, TableName),
		data.NameWithOwner,
		data.Name,
		data.Owner,
		data.URL,
		data.MainLangName,
		data.MainLangColor,
		data.Description,
		data.LastPushAt,
		data.CreatedAt,
		data.Stars,
		data.NameWithOwner,
	)
	if err != nil {
		lumber.Error(
			err,
			"Failed to update",
			TableName,
			"with most recent values for",
			data.NameWithOwner,
		)
	}
	lumber.Success("Updated", TableName, "with most recent values for", data.NameWithOwner)
}
