package dbwrapper

import (
	"database/sql"
)

type (
	Updatable interface {
		UpdateQuery() (string, []interface{})
		UpdateScanner(rows *sql.Rows) error
	}
)

func (db *DB) Update(item Updatable) error {
	query, args := item.UpdateQuery()
	return db.QueryAndScan(updatableScanner(item), query, args...)
}

func updatableScanner(item Updatable) Scanner {
	return func(row *sql.Rows) error {
		if err := item.UpdateScanner(row); err != nil {
			return err
		}
		return nil
	}
}
