package database

import (
	"database/sql"
	"net/url"
	"time"

	_ "github.com/microsoft/go-mssqldb"
	"maxintec-digitaltwin-api/internal/config"
)

func NewMSSQLConnection(cfg *config.Config) (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", "WiserSeDb-MAXI")
	query.Add("encrypt", "disable")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:     cfg.DBHost + ":1433",
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}