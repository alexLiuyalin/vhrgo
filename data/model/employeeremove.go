package model

import (
	"database/sql"
)

type EmployeeRemove struct {
	ID         int64          `json:"id" db:"id"`
	EmployeeID sql.NullInt64  `json:"eid" db:"eid"`               // 员工编号
	AfterDepID sql.NullInt64  `json:"afterDepId" db:"afterDepId"` // 调动后部门
	AfterJobID sql.NullInt64  `json:"afterJobId" db:"afterJobId"` // 调动后职位
	RemoveDate sql.NullTime   `json:"removeDate" db:"removeDate"` // 调动日期
	Reason     sql.NullString `json:"reason" db:"reason"`         // 调动原因
	Remark     sql.NullString `json:"remark" db:"remark"`
}
