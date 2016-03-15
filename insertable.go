package dbwrapper

import (
	"database/sql"
)

type (
	Insertable interface {
		InsertQuery() (string, []interface{})
		InsertScanner(rows *sql.Rows) error
	}
)

func (db *DB) Insert(item Insertable) error {
	query, args := item.InsertQuery()
	return db.QueryAndScan(insertableScanner(item), query, args...)
}

func insertableScanner(item Insertable) Scanner {
	return func(row *sql.Rows) error {
		if err := item.InsertScanner(row); err != nil {
			return err
		}
		return nil
	}
}
