package db

import (
	"fmt"

	"github.com/Matt-Gleich/lumber"
)

// Delete and create a given table
func SafeCreate(createQuery string, tableName string) {
	if tableExists(tableName) {
		lumber.Info(tableName, "already exists")
		return
	}
	_, err := DB.Exec(createQuery)
	lumber.Fatal(err, "Failed to create table", tableName)
	lumber.Info("Created table", tableName)
}

// Reset a given table via a truncate
func ResetTable(tableName string) {
	_, err := DB.Exec(fmt.Sprintf("TRUNCATE TABLE %v;", tableName))
	lumber.Fatal(err, "Failed to truncate table", tableName)
}

// Checking if the table exists
func tableExists(tName string) bool {
	var (
		exists bool
		query  = fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = '%v');", tName)
	)
	err := DB.QueryRow(query).Scan(&exists)
	lumber.Fatal(err, "Failed to check if table exists")
	return exists
}
