package model

import "database/sql"

type Role struct {
	ID     int64          `json:"id" db:"id"`
	Name   sql.NullString `json:"name" db:"name"`
	NameZh sql.NullString `json:"nameZh" db:"nameZh"` // 角色名称
}
