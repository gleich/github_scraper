package account

import (
	"fmt"
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
	contributors int,
	hireable bool,
	location varchar(255),
	organizations int,
	website varchar(255),
	company varchar(255),
	bio varchar(255)
);`, TableName,
)
