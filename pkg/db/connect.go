package db

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gleich/lumber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbURL, err := url.Parse(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	password, _ := dbURL.User.Password()

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbURL.Hostname(),
		dbURL.User.Username(),
		password,
		strings.TrimPrefix(dbURL.Path, "/"),
		dbURL.Port(),
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	lumber.Success("Connected to database")
	return db, nil
}
