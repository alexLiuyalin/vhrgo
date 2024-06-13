package model

import "database/sql"

type EmpSalary struct {
	ID         int64         `json:"id" db:"id"`
	EmployeeID sql.NullInt64 `json:"eid" db:"eid"` // 员工编号
	SalaryID   sql.NullInt64 `json:"sid" db:"sid"`
}
