package dbwrapper

import (
	"database/sql"
)

type (
	Queryable interface {
		Load(row *sql.Rows) error
	}

	QueryableSlice interface {
		InitItem() Queryable
		Append(Queryable)
	}

	Result interface {
		Error() string
	}

	QueryFunc func(db *DB) Result
)

func (db *DB) FindAndLoadOne(item Queryable, query string, args ...interface{}) error {
	return db.QueryAndScan(itemScanner(item), query, args...)
}

func (db *DB) FindAndLoadMultiple(sl QueryableSlice, query string, args ...interface{}) error {
	return db.QueryAndScan(sliceScanner(sl), query, args...)
}

func itemScanner(item Queryable) Scanner {
	return func(row *sql.Rows) error {
		if err := item.Load(row); err != nil {
			return err
		}
		return nil
	}
}

func sliceScanner(sl QueryableSlice) Scanner {
	return func(row *sql.Rows) error {
		i := sl.InitItem()
		if err := i.Load(row); err != nil {
			return err
		}
		sl.Append(i)
		return nil
	}
}

func (db *DB) Find(qfunc QueryFunc) Result {
	return qfunc(db)
}
