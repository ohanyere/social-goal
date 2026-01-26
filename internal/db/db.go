package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addrs string, maxopenconns, maxidleconns int, maxidletime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addrs)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.SetMaxOpenConns(maxopenconns)
	db.SetMaxIdleConns(maxidleconns)

	duration, err := time.ParseDuration(maxidletime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables() {

}
