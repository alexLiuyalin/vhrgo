package model

import "database/sql"

type Department struct {
	ID       int64          `json:"id" db:"id"`
	Name     sql.NullString `json:"name" db:"name"` // 部门名称
	ParentID sql.NullInt64  `json:"parentId" db:"parentId"`
	DepPath  sql.NullString `json:"depPath" db:"depPath"`
	Enabled  sql.NullBool   `json:"enabled" db:"enabled"`
	IsParent sql.NullBool   `json:"isParent" db:"isParent"`
}
