package model

import (
	"database/sql"
)

type Position struct {
	ID         int64          `json:"id" db:"id"`
	Name       sql.NullString `json:"name" db:"name"`             // 职位
	CreateDate sql.NullTime   `json:"createDate" db:"createDate"` // 创建日期
	Enabled    sql.NullBool   `json:"enabled" db:"enabled"`
}
