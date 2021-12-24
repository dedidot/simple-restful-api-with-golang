package app

import (
	"database/sql"
	"time"
)

func NewDB() (*sql.DB) {
	db, err := sql.Open("mysql", "root:12345678@/pzn?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
