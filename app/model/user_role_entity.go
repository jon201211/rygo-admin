package model

import (
	//"rygo/app/db"

	//"xorm.io/core"
)

type UserRoleEntity struct {
	UserId int64 `json:"user_id" xorm:"not null pk comment('用户ID') BIGINT(20)"`
	RoleId int64 `json:"role_id" xorm:"not null pk comment('角色ID') BIGINT(20)"`
}
