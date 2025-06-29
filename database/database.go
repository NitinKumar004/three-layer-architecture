package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func Databasconnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:nitin@123@tcp(127.0.0.1:3306)/taskdatabse")
	if err != nil {
		return nil, errors.New("error to establishing the connections")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New("failed to connect the database")
	}
	return db, nil

}
