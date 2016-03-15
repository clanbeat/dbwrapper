package dbwrapper

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

type DB struct {
	conn *sql.DB
}

var dbUser = "clanbeat_api_user"
var dbPass = "teambeat"
var dbName = "clanbeat_development"
var dbHost = "localhost"
var dbPort = 5432

type Scanner func(row *sql.Rows) error

func ConnectDevelopment() *DB {
	return connect(fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPass))
}

func connect(dbConnString string) *DB {
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connecting to db")
	return &DB{conn: db}
}

func ConnectWithURL(URL string) *DB {
	dbConnString, _ := pq.ParseURL(URL)
	dbConnString += " sslmode=require"
	return connect(dbConnString)
}

func (db *DB) IsDown() bool {
	if _, err := db.conn.Exec("SELECT 1"); err != nil {
		log.Println(err)
		return true
	}
	return false
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

func (db *DB) query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *DB) QueryAndScan(scanner Scanner, query string, args ...interface{}) error {
	rows, err := db.query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		if err := scanner(rows); err != nil {
			return err
		}
	}
	return nil
}
