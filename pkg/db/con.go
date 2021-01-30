package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Matt-Gleich/lumber"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect to the database
func Connect() {
	dbURL, err := url.Parse(os.Getenv("DB_URL"))
	lumber.Error(err, "Failed to get postgresql url")
	password, _ := dbURL.User.Password()

	// Getting database info and validating it
	postgresInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbURL.Hostname(),
		dbURL.User.Username(),
		password,
		strings.TrimPrefix(dbURL.Path, "/"),
		dbURL.Port(),
	)
	db, err := sql.Open("postgres", postgresInfo)
	lumber.Error(err, "Failed to validate to postgres info:", postgresInfo)

	// Verifying connection to database
	err = db.Ping()
	lumber.Error(err, "Failed to connect to database")

	lumber.Info("Connected to database")
}
