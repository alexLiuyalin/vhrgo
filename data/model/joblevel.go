package model

import (
	"database/sql"
)

type JobLevel struct {
	ID         int64          `json:"id" db:"id"`
	Name       sql.NullString `json:"name" db:"name"`             // 职称名称
	TitleLevel sql.NullString `json:"titleLevel" db:"titleLevel"` // 职称级别
	CreateDate sql.NullTime   `json:"createDate" db:"createDate"` // 创建日期
	Enabled    sql.NullBool   `json:"enabled" db:"enabled"`
}
