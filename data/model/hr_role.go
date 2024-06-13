package model

import "database/sql"

type HRRole struct {
	ID     int64         `json:"id" db:"id"`
	HRID   sql.NullInt64 `json:"hrid" db:"hrid"`
	RoleID sql.NullInt64 `json:"rid" db:"rid"`
}
