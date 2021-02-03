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
	contributors int
);`, TableName,
)
