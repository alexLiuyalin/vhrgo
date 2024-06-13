package model

import "database/sql"

type HR struct {
	ID        int64          `json:"id" db:"id"`               // hrID
	Name      sql.NullString `json:"name" db:"name"`           // 姓名
	Phone     sql.NullString `json:"phone" db:"phone"`         // 手机号码
	Telephone sql.NullString `json:"telephone" db:"telephone"` // 住宅电话
	Address   sql.NullString `json:"address" db:"address"`     // 联系地址
	Enabled   sql.NullBool   `json:"enabled" db:"enabled"`
	Username  sql.NullString `json:"username" db:"username"` // 用户名
	Password  sql.NullString `json:"password" db:"password"` // 密码
	Userface  sql.NullString `json:"userface" db:"userface"`
	Remark    sql.NullString `json:"remark" db:"remark"`
}
