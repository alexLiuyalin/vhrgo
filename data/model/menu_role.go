package model

import "database/sql"

type MenuRole struct {
	ID  int64         `json:"id" db:"id"`
	MID sql.NullInt64 `json:"mid" db:"mid"`
	RID sql.NullInt64 `json:"rid" db:"rid"`
}
