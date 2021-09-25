package db

import (
	"fmt"

	"github.com/gleich/lumber/v2"
)

// Delete and create a given table
func HardResetTable(createQuery string, tableName string) {
	if tableExists(tableName) {
		_, err := DB.Exec(fmt.Sprintf("DROP TABLE %v;", tableName))
		if err != nil {
			lumber.Fatal(err, "Failed to delete table", tableName)
		}
		lumber.Info("Deleted table", tableName)
	}
	_, err := DB.Exec(createQuery)
	if err != nil {
		lumber.Fatal(err, "Failed to create table", tableName, "\n", createQuery)
	}
	lumber.Info("Created table", tableName)
}

// Checking if the table exists
func tableExists(tName string) bool {
	var (
		exists bool
		query  = fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = '%s');", tName)
	)
	err := DB.QueryRow(query).Scan(&exists)
	if err != nil {
		lumber.Fatal(err, "Failed to check if table exists")
	}
	return exists
}
