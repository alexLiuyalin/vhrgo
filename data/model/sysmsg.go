package model

import "database/sql"

type SysMsg struct {
	ID    int64         `json:"id" db:"id"`
	MID   sql.NullInt64 `json:"mid" db:"mid"`     // 消息id
	Type  sql.NullInt64 `json:"type" db:"type"`   // 0表示群发消息
	HRID  sql.NullInt64 `json:"hrid" db:"hrid"`   // 这条消息是给谁的
	State sql.NullInt64 `json:"state" db:"state"` // 0 未读 1 已读
}
