package dbwrapper

import (
	"database/sql"
)

func NullID(idValue int64) sql.NullInt64 {
	if idValue == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: idValue, Valid: true}
}

func NullString(value string) sql.NullString {
	if len(value) == 0 {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: value, Valid: true}
}
