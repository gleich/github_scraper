package db

import (
	"fmt"

	"github.com/Matt-Gleich/lumber"
)

// Delete and create a given table
func HardResetTable(createQuery string, tableName string) {
	if tableExists(tableName) {
		_, err := DB.Exec(fmt.Sprintf("DROP TABLE %v;", tableName))
		lumber.Fatal(err, "Failed to delete table", tableName)
		lumber.Info("Deleted table", tableName)
	}
	_, err := DB.Exec(createQuery)
	lumber.Fatal(err, "Failed to create table", tableName, "\n", createQuery)
	lumber.Info("Created table", tableName)
}

// Checking if the table exists
func tableExists(tName string) bool {
	var (
		exists bool
		query  = fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = '%s');", tName)
	)
	err := DB.QueryRow(query).Scan(&exists)
	lumber.Fatal(err, "Failed to check if table exists")
	return exists
}
