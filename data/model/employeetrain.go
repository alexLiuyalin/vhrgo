package model

import (
	"database/sql"
)

type EmployeeTrain struct {
	ID           int64          `json:"id" db:"id"`
	EmployeeID   sql.NullInt64  `json:"eid" db:"eid"`                   // 员工编号
	TrainDate    sql.NullTime   `json:"trainDate" db:"trainDate"`       // 培训日期
	TrainContent sql.NullString `json:"trainContent" db:"trainContent"` // 培训内容
	Remark       sql.NullString `json:"remark" db:"remark"`             // 备注
}
