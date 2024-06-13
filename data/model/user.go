package model

import (
	"database/sql"
)

type User struct {
	ID       int64          `json:"id" db:"id"`             // hrID
	Status   sql.NullBool   `json:"status" db:"status"`     // 状态
	Username sql.NullString `json:"username" db:"username"` // 用户名
	Password sql.NullString `json:"password" db:"password"` // 密码
}
