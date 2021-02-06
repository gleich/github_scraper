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
	pname varchar(255),
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
			pname,
			url,
			main_lang_name,
			main_lang_color,
			description,
			stars) VALUES (
				'%s', '%s', '%s', '%s', '%s', %v
			) ;
		`,
		TableName,
		data.Name,
		data.URL,
		data.MainLangName,
		data.MainLangColor,
		data.Description,
		data.Stars,
	))
	lumber.Error(err, "Failed to update", TableName, "with latest information")
	lumber.Success("Updated", TableName, "with latest data")
}
