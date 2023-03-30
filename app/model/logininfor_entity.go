package model

import (
	"time"
)

type LogininforEntity struct {
	InfoId        int64     `json:"info_id" xorm:"not null pk autoincr comment('访问ID') BIGINT(20)"`
	LoginName     string    `json:"login_name" xorm:"default '' comment('登录账号') VARCHAR(50)"`
	Ipaddr        string    `json:"ipaddr" xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation string    `json:"login_location" xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser       string    `json:"browser" xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os            string    `json:"os" xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status        string    `json:"status" xorm:"default '0' comment('登录状态（0成功 1失败）') CHAR(1)"`
	Msg           string    `json:"msg" xorm:"default '' comment('提示消息') VARCHAR(255)"`
	LoginTime     time.Time `json:"login_time" xorm:"comment('访问时间') DATETIME"`
}

// Fill with you ideas below.
//查询列表请求参数
type LogininforSelectPageReq struct {
	LoginName     string `form:"loginName"`     //登陆名
	Status        string `form:"status"`        //状态
	Ipaddr        string `form:"ipaddr"`        //登录地址
	BeginTime     string `form:"beginTime"`     //数据范围
	EndTime       string `form:"endTime"`       //开始时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}
