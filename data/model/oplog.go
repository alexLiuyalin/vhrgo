package model

import (
	"database/sql"
)

type OpLog struct {
	ID      int64          `json:"id" db:"id"`
	AddDate sql.NullTime   `json:"addDate" db:"addDate"` // 添加日期
	Operate sql.NullString `json:"operate" db:"operate"` // 操作内容
	HRID    sql.NullInt64  `json:"hrid" db:"hrid"`       // 操作员ID
}
