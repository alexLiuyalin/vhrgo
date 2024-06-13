package model

import (
	"database/sql"
)

type Salary struct {
	ID                   int64           `json:"id" db:"id"`
	BasicSalary          sql.NullInt64   `json:"basicSalary" db:"basicSalary"`                   // 基本工资
	Bonus                sql.NullInt64   `json:"bonus" db:"bonus"`                               // 奖金
	LunchSalary          sql.NullInt64   `json:"lunchSalary" db:"lunchSalary"`                   // 午餐补助
	TrafficSalary        sql.NullInt64   `json:"trafficSalary" db:"trafficSalary"`               // 交通补助
	AllSalary            sql.NullInt64   `json:"allSalary" db:"allSalary"`                       // 应发工资
	PensionBase          sql.NullInt64   `json:"pensionBase" db:"pensionBase"`                   // 养老金基数
	PensionPer           sql.NullFloat64 `json:"pensionPer" db:"pensionPer"`                     // 养老金比率
	CreateDate           sql.NullTime    `json:"createDate" db:"createDate"`                     // 启用时间
	MedicalBase          sql.NullInt64   `json:"medicalBase" db:"medicalBase"`                   // 医疗基数
	MedicalPer           sql.NullFloat64 `json:"medicalPer" db:"medicalPer"`                     // 医疗保险比率
	AccumulationFundBase sql.NullInt64   `json:"accumulationFundBase" db:"accumulationFundBase"` // 公积金基数
	AccumulationFundPer  sql.NullFloat64 `json:"accumulationFundPer" db:"accumulationFundPer"`   // 公积金比率
	Name                 sql.NullString  `json:"name" db:"name"`
}
