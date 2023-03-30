package model

type RoleDeptEntity struct {
	RoleId int64 `json:"role_id" xorm:"not null pk comment('角色ID') BIGINT(20)"`
	DeptId int64 `json:"dept_id" xorm:"not null pk comment('部门ID') BIGINT(20)"`
}
