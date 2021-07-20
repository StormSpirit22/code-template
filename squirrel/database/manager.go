package database

import (
	"db_connect/squirrel/util/postgres"
	"github.com/jmoiron/sqlx"

)

type Manager struct {
	db *sqlx.DB
}

func NewManager() (*Manager, error) {
	mgr := &Manager{}

	cfg, err := postgres.NewPgSqlConfig("postgres")
	if err != nil {
		return nil, err
	}

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		return nil, err
	}
	mgr.db = db
	err = mgr.createUserTable()
	if err != nil {
		return nil, err
	}
	return mgr, nil
}
