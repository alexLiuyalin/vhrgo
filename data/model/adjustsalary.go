package model

import (
	"database/sql"
)

type AdjustSalary struct {
	ID           int64          `json:"id" db:"id"`
	EmployeeID   sql.NullInt64  `json:"eid" db:"eid"`
	AsDate       sql.NullTime   `json:"asDate" db:"asDate"`             // 调薪日期
	BeforeSalary sql.NullInt64  `json:"beforeSalary" db:"beforeSalary"` // 调前薪资
	AfterSalary  sql.NullInt64  `json:"afterSalary" db:"afterSalary"`   // 调后薪资
	Reason       sql.NullString `json:"reason" db:"reason"`             // 调薪原因
	Remark       sql.NullString `json:"remark" db:"remark"`             // 备注
}
