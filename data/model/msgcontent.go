package model

import (
	"database/sql"
	"time"
)

type MsgContent struct {
	ID         int64          `json:"id" db:"id"`
	Title      sql.NullString `json:"title" db:"title"`
	Message    sql.NullString `json:"message" db:"message"`
	CreateDate time.Time      `json:"createDate" db:"createDate"`
}
