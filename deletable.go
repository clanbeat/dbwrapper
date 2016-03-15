package dbwrapper

type (
	Deletable interface {
		DeleteQuery() (string, []interface{})
	}
)

func (db *DB) Delete(item Deletable) error {
	query, args := item.DeleteQuery()
	_, err := db.exec(query, args...)
	return err
}
