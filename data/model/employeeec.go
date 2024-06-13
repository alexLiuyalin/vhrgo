package model

import (
	"database/sql"
)

type EmployeeEC struct {
	ID         int64          `json:"id" db:"id"`
	EmployeeID sql.NullInt64  `json:"eid" db:"eid"`           // 员工编号
	ECDate     sql.NullTime   `json:"ecDate" db:"ecDate"`     // 奖罚日期
	ECReason   sql.NullString `json:"ecReason" db:"ecReason"` // 奖罚原因
	ECPoint    sql.NullInt64  `json:"ecPoint" db:"ecPoint"`   // 奖罚分
	ECType     sql.NullInt64  `json:"ecType" db:"ecType"`     // 奖罚类别，0：奖，1：罚
	Remark     sql.NullString `json:"remark" db:"remark"`     // 备注
}
