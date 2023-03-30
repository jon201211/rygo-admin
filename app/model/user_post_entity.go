package model

import (
	//"rygo/app/db"

	//"xorm.io/core"
)

type UserPostEntity struct {
	UserId int64 `json:"user_id" xorm:"not null pk comment('用户ID') BIGINT(20)"`
	PostId int64 `json:"post_id" xorm:"not null pk comment('岗位ID') BIGINT(20)"`
}
