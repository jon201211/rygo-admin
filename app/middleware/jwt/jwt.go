package jwt

import (
	"net/http"
	"rygo/app/token"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		// 按空格分割
		tokenStr := ""
		if authHeader != "" { //1先从header取
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) > 1 {
				tokenStr = parts[1]
			}
		} else {
			tokenStr = ctx.Request.Header.Get("token")
		}

		if tokenStr == "" { //2从url取
			tokenStr = ctx.Query("token")
		}
		cookie, err := ctx.Request.Cookie("token")
		if err == nil { //3 从cookie取
			tokenStr = cookie.Value
		}

		if tokenStr == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取token失败",
			})
			ctx.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		t, err := token.VerifyAuthToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			ctx.Abort()
			return
		}
		// 将当前请求的uid信息保存到请求的上下文c上
		ctx.Set("tenantId", t.Claim.TenantId)
		ctx.Set("userId", t.Claim.UserId)
		ctx.Set("loginName", t.Claim.Subject)
		//判断是否有新的token生成
		if t.NewToken != "" {
			ctx.Writer.Header().Set("nt", t.NewToken)
		}
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
