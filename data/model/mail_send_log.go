package model

import (
	"database/sql"
)

type MailSendLog struct {
	MsgID      sql.NullString `json:"msgId" db:"msgId"`
	EmpID      sql.NullInt64  `json:"empId" db:"empId"`
	Status     sql.NullInt64  `json:"status" db:"status"` // 0发送中，1发送成功，2发送失败
	RouteKey   sql.NullString `json:"routeKey" db:"routeKey"`
	Exchange   sql.NullString `json:"exchange" db:"exchange"`
	Count      sql.NullInt64  `json:"count" db:"count"`     // 重试次数
	TryTime    sql.NullTime   `json:"tryTime" db:"tryTime"` // 第一次重试时间
	CreateTime sql.NullTime   `json:"createTime" db:"createTime"`
	UpdateTime sql.NullTime   `json:"updateTime" db:"updateTime"`
}
