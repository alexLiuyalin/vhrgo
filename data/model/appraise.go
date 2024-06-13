package model

import (
	"database/sql"
)

type Appraise struct {
	ID         int64          `json:"id" db:"id"`
	EmployeeID sql.NullInt64  `json:"eid" db:"eid"`
	AppDate    sql.NullTime   `json:"appDate" db:"appDate"`       // 考评日期
	AppResult  sql.NullString `json:"appResult" db:"appResult"`   // 考评结果
	AppContent sql.NullString `json:"appContent" db:"appContent"` // 考评内容
	Remark     sql.NullString `json:"remark" db:"remark"`         // 备注
}
