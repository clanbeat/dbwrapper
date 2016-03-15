package dbwrapper

import (
	"database/sql"
	"fmt"
)

func (db *DB) Transact(txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	return txFunc(tx)
}
