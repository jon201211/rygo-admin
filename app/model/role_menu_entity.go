package model

type RoleMenuEntity struct {
	RoleId int64 `json:"role_id" xorm:"not null pk comment('角色ID') BIGINT(20)"`
	MenuId int64 `json:"menu_id" xorm:"not null pk comment('菜单ID') BIGINT(20)"`
}
