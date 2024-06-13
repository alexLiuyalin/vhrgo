package model

import (
	"database/sql"
)

type Employee struct {
	ID             int64           `json:"id" db:"id"`                         // 员工编号
	Name           sql.NullString  `json:"name" db:"name"`                     // 员工姓名
	Gender         sql.NullString  `json:"gender" db:"gender"`                 // 性别
	Birthday       sql.NullTime    `json:"birthday" db:"birthday"`             // 出生日期
	IDCard         sql.NullString  `json:"idCard" db:"idCard"`                 // 身份证号
	Wedlock        sql.NullString  `json:"wedlock" db:"wedlock"`               // 婚姻状况
	NationID       sql.NullInt64   `json:"nationId" db:"nationId"`             // 民族
	NativePlace    sql.NullString  `json:"nativePlace" db:"nativePlace"`       // 籍贯
	PoliticID      sql.NullInt64   `json:"politicId" db:"politicId"`           // 政治面貌
	Email          sql.NullString  `json:"email" db:"email"`                   // 邮箱
	Phone          sql.NullString  `json:"phone" db:"phone"`                   // 电话号码
	Address        sql.NullString  `json:"address" db:"address"`               // 联系地址
	DepartmentID   sql.NullInt64   `json:"departmentId" db:"departmentId"`     // 所属部门
	JobLevelID     sql.NullInt64   `json:"jobLevelId" db:"jobLevelId"`         // 职称ID
	PosID          sql.NullInt64   `json:"posId" db:"posId"`                   // 职位ID
	EngageForm     sql.NullString  `json:"engageForm" db:"engageForm"`         // 聘用形式
	TiptopDegree   sql.NullString  `json:"tiptopDegree" db:"tiptopDegree"`     // 最高学历
	Specialty      sql.NullString  `json:"specialty" db:"specialty"`           // 所属专业
	School         sql.NullString  `json:"school" db:"school"`                 // 毕业院校
	BeginDate      sql.NullTime    `json:"beginDate" db:"beginDate"`           // 入职日期
	WorkState      sql.NullString  `json:"workState" db:"workState"`           // 在职状态
	WorkID         sql.NullString  `json:"workID" db:"workID"`                 // 工号
	ContractTerm   sql.NullFloat64 `json:"contractTerm" db:"contractTerm"`     // 合同期限
	ConversionTime sql.NullTime    `json:"conversionTime" db:"conversionTime"` // 转正日期
	NotWorkDate    sql.NullTime    `json:"notWorkDate" db:"notWorkDate"`       // 离职日期
	BeginContract  sql.NullTime    `json:"beginContract" db:"beginContract"`   // 合同起始日期
	EndContract    sql.NullTime    `json:"endContract" db:"endContract"`       // 合同终止日期
	WorkAge        sql.NullInt64   `json:"workAge" db:"workAge"`               // 工龄
}
