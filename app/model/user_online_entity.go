package model

import (
	"time"
)

type UserOnline struct {
	Sessionid      string    `json:"sessionId"  xorm:"not null default '' comment('用户会话id') VARCHAR(250)"`
	Token          string    `json:"token"      xorm:"default '' comment('jwt token') VARCHAR(255)"`
	LoginName      string    `json:"login_name" xorm:"default '' comment('登录账号') VARCHAR(50)"`
	DeptName       string    `json:"dept_name"  xorm:"default '' comment('部门名称') VARCHAR(50)"`
	Ipaddr         string    `json:"ipaddr" xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation  string    `json:"login_location" xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser        string    `json:"browser" xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os             string    `json:"os" xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status         string    `json:"status" xorm:"default '' comment('在线状态on_line在线off_line离线') VARCHAR(10)"`
	StartTimestamp time.Time `json:"start_timestamp" xorm:"comment('session创建时间') DATETIME"`
	LastAccessTime time.Time `json:"last_access_time" xorm:"comment('session最后访问时间') DATETIME"`
	ExpireTime     int       `json:"expire_time" xorm:"default 0 comment('超时时间，单位为分钟') INT(5)"`
}

//新增页面请求参数
type UserOnlineAddReq struct {
	LoginName      string    `form:"loginName" binding:"required"`
	DeptName       string    `form:"deptName" binding:"required"`
	Ipaddr         string    `form:"ipaddr" `
	LoginLocation  string    `form:"loginLocation" `
	Browser        string    `form:"browser" `
	Os             string    `form:"os" `
	Status         string    `form:"status" binding:"required"`
	StartTimestamp time.Time `form:"startTimestamp" `
	LastAccessTime time.Time `form:"lastAccessTime" `
	ExpireTime     int       `form:"expireTime" `
}

//修改页面请求参数
type UserOnlineEditReq struct {
	SessionId      string    `form:"sessionId" binding:"required"`
	LoginName      string    `form:"loginName" binding:"required"`
	DeptName       string    `form:"deptName" binding:"required"`
	Ipaddr         string    `form:"ipaddr" `
	LoginLocation  string    `form:"loginLocation" `
	Browser        string    `form:"browser" `
	Os             string    `form:"os" `
	Status         string    `form:"status" binding:"required"`
	StartTimestamp time.Time `form:"startTimestamp" `
	LastAccessTime time.Time `form:"lastAccessTime" `
	ExpireTime     int       `form:"expireTime" `
}

//分页请求参数
type UserOnlineSelectPageReq struct {
	SessionId      string    `form:"sessionId"`      //用户会话id
	LoginName      string    `form:"loginName"`      //登录账号
	DeptName       string    `form:"deptName"`       //部门名称
	Ipaddr         string    `form:"ipaddr"`         //登录IP地址
	LoginLocation  string    `form:"loginLocation"`  //登录地点
	Browser        string    `form:"browser"`        //浏览器类型
	Os             string    `form:"os"`             //操作系统
	Status         string    `form:"status"`         //在线状态on_line在线off_line离线
	StartTimestamp time.Time `form:"startTimestamp"` //session创建时间
	LastAccessTime time.Time `form:"lastAccessTime"` //session最后访问时间
	ExpireTime     int       `form:"expireTime"`     //超时时间，单位为分钟
	BeginTime      string    `form:"beginTime"`      //开始时间
	EndTime        string    `form:"endTime"`        //结束时间
	PageNum        int       `form:"pageNum"`        //当前页码
	PageSize       int       `form:"pageSize"`       //每页数
	OrderByColumn  string    `form:"orderByColumn"`  //排序字段
	IsAsc          string    `form:"isAsc"`          //排序方式
}
