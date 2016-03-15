package dbwrapper

type (
	Paranoid interface {
		SetDeleteQuery() (string, []interface{})
	}
)

func (db *DB) ParanoidDelete(item Paranoid) error {
	query, args := item.SetDeleteQuery()
	_, err := db.exec(query, args...)
	return err
}
