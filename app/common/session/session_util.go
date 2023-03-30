package session

import (
	"rygo/app/dao"
	"rygo/app/global"
	"rygo/app/model"

	"sync"

	"github.com/gin-gonic/gin"
)

//用户session列表
var SessionList sync.Map

//判断是否是系统管理员
func IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

func IsAdminUser(user *model.SysUser) bool {
	if user.UserId == 1 {
		return true
	} else {
		return false
	}
}

// 判断用户是否已经登录
func IsSignedIn(ctx *gin.Context) bool {
	_, exist := ctx.Get(global.USER_ID)
	if exist {
		return true
	}
	return false
}

// 获得用户信息详情
func GetProfile(ctx *gin.Context) *model.SysUser {
	userId, exist := ctx.Get(global.USER_ID)
	if exist == false {
		return nil
	}
	user := model.SysUser{}
	user.UserId = userId.(int64)
	_, err := dao.UserDao.FindOne(&user)
	if err != nil {
		return nil
	}
	//err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		return nil
	}
	return &user
}

// 获得用户信息详情
func GetTenantId(ctx *gin.Context) int64 {
	u := GetProfile(ctx)
	return u.TenantId
}
