package model

import (
	"database/sql"
)

type Employee struct {
	ID             int64          `json:"id" db:"id"`
	User           User           `json:"user" db:"user"`
	UserId         int64          `json:"user_id" db:"user_id"`                 // 用户id
	WorkId         sql.NullString `json:"work_id" db:"work_id"`                 // 工号
	Name           sql.NullString `json:"name" db:"name"`                       // 员工姓名
	Gender         sql.NullString `json:"gender" db:"gender"`                   // 性别
	Picture        sql.NullString `json:"picture" db:"picture"`                 // 头像
	IdCard         sql.NullString `json:"id_card" db:"id_card"`                 // 身份证号
	IdCardPicture  sql.NullString `json:"id_card_picture" db:"id_card_picture"` // 身份证图片
	Wedlock        sql.NullString `json:"wedlock" db:"wedlock"`                 // 婚姻状况
	NationId       sql.NullInt64  `json:"nation_id" db:"nation_id"`             // 民族
	NativePlace    sql.NullString `json:"native_place" db:"native_place"`       // 籍贯
	PoliticId      sql.NullInt64  `json:"politic_id" db:"politic_id"`           // 政治面貌
	Email          sql.NullString `json:"email" db:"email"`                     // 邮箱
	Phone          sql.NullString `json:"phone" db:"phone"`                     // 电话号码
	Address        sql.NullString `json:"address" db:"address"`                 // 联系地址
	DepartmentId   sql.NullInt64  `json:"department_id" db:"department_id"`     // 所属部门
	JobLevelId     sql.NullInt64  `json:"job_level_id" db:"job_level_id"`       // 职称ID
	PosId          sql.NullInt64  `json:"pos_id" db:"pos_id"`                   // 职位ID
	EngageForm     sql.NullString `json:"engage_form" db:"engage_form"`         // 聘用形式
	TiptopDegree   sql.NullString `json:"tiptop_degree" db:"tiptop_degree"`     // 最高学历
	Specialty      sql.NullString `json:"specialty" db:"specialty"`             // 所属专业
	School         sql.NullString `json:"school" db:"school"`                   // 毕业院校
	BeginDate      sql.NullTime   `json:"begin_date" db:"begin_date"`           // 入职日期
	NotWorkDate    sql.NullTime   `json:"not_work_date" db:"not_work_date"`     // 离职日期
	WorkStatus     sql.NullInt32  `json:"work_status" db:"work_status"`         // 在职状态
	ContractTerm   sql.NullInt32  `json:"contract_term" db:"contract_term"`     // 合同期限
	ConversionTime sql.NullTime   `json:"conversion_time" db:"conversion_time"` // 转正日期
	BeginContract  sql.NullTime   `json:"begin_contract" db:"begin_contract"`   // 合同起始日期
	EndContract    sql.NullTime   `json:"end_contract" db:"end_contract"`       // 合同终止日期
	WorkAge        sql.NullInt32  `json:"work_age" db:"work_age"`               // 工龄
}
