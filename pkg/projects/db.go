package projects

import "fmt"

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
