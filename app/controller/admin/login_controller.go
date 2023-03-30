package admin

import (
	"encoding/json"
	"net/http"
	"rygo/app/controller/base"
	"rygo/app/global"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/token"
	"rygo/app/utils/gconv"
	"rygo/app/utils/ip"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
)

type RegisterReq struct {
	UserName     string `form:"username"  binding:"required,min=4,max=30"`
	Password     string `form:"password" binding:"required,min=6,max=30"`
	ValidateCode string `form:"validateCode" binding:"required,min=4,max=10"`
	IdKey        string `form:"idkey" binding:"required,min=5,max=30"`
}
type LoginController struct {
	base.BaseController
}

// 登陆页面
func (c *LoginController) Login(ctx *gin.Context) {

	if strings.EqualFold(ctx.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		response.ErrorResp(ctx).SetMsg("未登录或登录超时。请重新登录").WriteJsonExit()
		return
	}

	response.BuildTpl(ctx, "login").WriteTpl()
}

// 图形验证码
func (c *LoginController) CaptchaImage(ctx *gin.Context) {
	//config struct for digits
	//数字验证码配置
	//var configD = base64Captcha.ConfigDigit{
	//	Height:     80,
	//	Width:      240,
	//	MaxSkew:    0.7,
	//	DotCount:   80,
	//	CaptchaLen: 5,
	//}
	//config struct for audio
	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	ctx.JSON(http.StatusOK, model.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

//验证登陆
func (c *LoginController) CheckLogin(ctx *gin.Context) {
	var req = RegisterReq{}
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	errTimes := service.LogininforService.GetPasswordCounts(req.UserName)
	if errTimes > 5 { //超过5次错误开始校验验证码
		//比对验证码
		verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)
		if !verifyResult {
			response.ErrorResp(ctx).SetMsg("验证码不正确").WriteJsonExit()
			return
		}
	}
	isLock := service.LogininforService.CheckLock(req.UserName)
	if isLock {
		response.ErrorResp(ctx).SetMsg("账号已锁定，请30分钟后再试").WriteJsonExit()
		return
	}
	//验证账号密码
	user, err := service.UserService.SignIn(req.UserName, req.Password)
	if err != nil {
		errTimes := service.LogininforService.SetPasswordCounts(req.UserName)
		having := 10 - errTimes
		c.SaveLogs(ctx, &req, "账号或密码不正确") //记录日志
		response.ErrorResp(ctx).SetMsg("账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定").WriteJsonExit()
	} else {
		//保存在线状态
		cookie, _ := ctx.Request.Cookie("token")
		token, _ := token.New(user.LoginName, user.UserId, user.TenantId).CreateToken()
		if cookie == nil {
			cookie = &http.Cookie{
				Name:     "token",
				Value:    token,
				HttpOnly: true,
			}
			http.SetCookie(ctx.Writer, cookie)
		}
		ctx.SetCookie(cookie.Name, token, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.SameSite, cookie.Secure, cookie.HttpOnly)
		// 生成token
		c.SaveUserToSession(user, ctx)
		c.SaveLogs(ctx, &req, "登陆成功") //记录日志
		response.SucessResp(ctx).SetData(token).SetMsg("登陆成功").WriteJsonExit()
	}
}

func (c *LoginController) SaveLogs(ctx *gin.Context, req *RegisterReq, msg string) {
	var logininfor model.LogininforEntity
	logininfor.LoginName = req.UserName
	logininfor.Ipaddr = ctx.ClientIP()
	userAgent := ctx.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	logininfor.Os = ua.OS()
	logininfor.Browser, _ = ua.Browser()
	logininfor.LoginTime = time.Now()
	logininfor.LoginLocation = ip.GetCityByIp(logininfor.Ipaddr)
	logininfor.Msg = msg
	logininfor.Status = "0"
	service.LogininforService.Insert(&logininfor)
}

//保存用户信息到session
func (c *LoginController) SaveUserToSession(user *model.SysUser, ctx *gin.Context) {
	tmp, _ := json.Marshal(user)
	sessionId := user.UserId
	global.SessionList.Store(sessionId, string(tmp))
	//save to db
	userAgent := ctx.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	loginIp := ctx.ClientIP()
	loginLocation := ip.GetCityByIp(loginIp)
	//移除登陆次数记录
	service.LogininforService.RemovePasswordCounts(user.UserName)
	//
	var userOnline model.UserOnline
	userOnline.LoginName = user.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.DeptName = ""
	userOnline.Ipaddr = loginIp
	userOnline.ExpireTime = 1440
	userOnline.StartTimestamp = time.Now()
	userOnline.LastAccessTime = time.Now()
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	service.OnlineService.Delete(&userOnline)
	service.OnlineService.Insert(&userOnline)
}
