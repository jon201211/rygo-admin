package admin

import (
	"net/http"
	"os"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	base.BaseController
}

//用户资料页面
func (c *ProfileController) Profile(ctx *gin.Context) {
	user := service.UserService.GetProfile(ctx)
	response.BuildTpl(ctx, "system/user/profile/profile").WriteTpl(gin.H{
		"user": user,
	})
}

//修改用户信息
func (c *ProfileController) Update(ctx *gin.Context) {
	var req *model.UserProfileReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}

	err := service.UserService.UpdateProfile(req, ctx)

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("用户管理", req).WriteJsonExit()
	}
}

//修改用户密码
func (c *ProfileController) UpdatePassword(ctx *gin.Context) {
	var req *model.UserPasswordReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	}

	err := service.UserService.UpdatePassword(req, ctx)

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("修改用户密码", req).WriteJsonExit()
	}
}

//修改头像页面
func (c *ProfileController) Avatar(ctx *gin.Context) {
	user := service.UserService.GetProfile(ctx)
	response.BuildTpl(ctx, "system/user/profile/avatar").WriteTpl(gin.H{
		"user": user,
	})
}

//修改密码页面
func (c *ProfileController) EditPwd(ctx *gin.Context) {
	user := service.UserService.GetProfile(ctx)
	response.BuildTpl(ctx, "system/user/profile/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}

//检查登陆名是否存在
func (c *ProfileController) CheckLoginNameUnique(ctx *gin.Context) {
	var req *model.UserCheckLoginNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.UserService.CheckLoginName(req.LoginName)

	if result {
		ctx.Writer.WriteString("1")
	} else {
		ctx.Writer.WriteString("0")
	}
}

//检查邮箱是否存在
func (c *ProfileController) CheckEmailUnique(ctx *gin.Context) {
	var req *model.UserCheckEmailReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.UserService.CheckEmailUnique(req.UserId, req.Email)

	if result {
		ctx.Writer.WriteString("1")
	} else {
		ctx.Writer.WriteString("0")
	}
}

//检查邮箱是否存在
func (c *ProfileController) CheckEmailUniqueAll(ctx *gin.Context) {
	var req *model.UserCheckEmailAllReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.UserService.CheckEmailUniqueAll(req.Email)

	if result {
		ctx.Writer.WriteString("1")
	} else {
		ctx.Writer.WriteString("0")
	}
}

//检查手机号是否存在
func (c *ProfileController) CheckPhoneUnique(ctx *gin.Context) {
	var req *model.UserCheckPhoneReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.UserService.CheckPhoneUnique(req.UserId, req.Phonenumber)

	if result {
		ctx.Writer.WriteString("1")
	} else {
		ctx.Writer.WriteString("0")
	}

}

//检查手机号是否存在
func (c *ProfileController) CheckPhoneUniqueAll(ctx *gin.Context) {
	var req *model.UserCheckPhoneAllReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	result := service.UserService.CheckPhoneUniqueAll(req.Phonenumber)

	if result {
		ctx.Writer.WriteString("1")
	} else {
		ctx.Writer.WriteString("0")
	}

}

//校验密码是否正确
func (c *ProfileController) CheckPassword(ctx *gin.Context) {
	var req *model.UserCheckPasswordReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	user := service.UserService.GetProfile(ctx)

	result := service.UserService.CheckPassword(user, req.Password)

	if result {
		ctx.Writer.WriteString("true")
	} else {
		ctx.Writer.WriteString("false")
	}
}

//保存头像
func (c *ProfileController) UpdateAvatar(ctx *gin.Context) {
	user := service.UserService.GetProfile(ctx)

	curDir, err := os.Getwd()

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	saveDir := curDir + "/public/upload/"

	fileHead, err := ctx.FormFile("avatarfile")

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("没有获取到上传文件").Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	curdate := time.Now().UnixNano()
	filename := user.LoginName + strconv.FormatInt(curdate, 10) + ".png"
	dts := saveDir + filename

	if err := ctx.SaveUploadedFile(fileHead, dts); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}

	avatar := "/upload/" + filename

	err = service.UserService.UpdateAvatar(avatar, ctx)

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("保存头像", gin.H{"userid": user.UserId}).WriteJsonExit()
	}
}
