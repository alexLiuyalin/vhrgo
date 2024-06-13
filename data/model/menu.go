package model

import "database/sql"

type Menu struct {
	ID          int64          `json:"id" db:"id"`
	URL         sql.NullString `json:"url" db:"url"`
	Path        sql.NullString `json:"path" db:"path"`
	Component   sql.NullString `json:"component" db:"component"`
	Name        sql.NullString `json:"name" db:"name"`
	IconCls     sql.NullString `json:"iconCls" db:"iconCls"`
	KeepAlive   sql.NullBool   `json:"keepAlive" db:"keepAlive"`
	RequireAuth sql.NullBool   `json:"requireAuth" db:"requireAuth"`
	ParentID    sql.NullInt64  `json:"parentId" db:"parentId"`
	Enabled     sql.NullBool   `json:"enabled" db:"enabled"`
}
