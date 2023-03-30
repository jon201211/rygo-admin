package model

import (
	"time"
)

// 数据表映射结构体
type SysTenant struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr comment('id') bigint"`
	DelFlag    string    `json:"del_flag" xorm:"comment('删除标志（0代表存在 2代表删除）') char(1)"`
	CreateBy   string    `json:"create_by" xorm:"comment('创建者') varchar(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') datetime"`
	UpdateBy   string    `json:"update_by" xorm:"comment('更新者') varchar(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') datetime"`
	Name       string    `json:"name" xorm:"comment('商户名称') varchar(32)"`
	Address    string    `json:"address" xorm:"comment('联系地址') varchar(64)"`
	Manager    string    `json:"manager" xorm:"comment('负责人') varchar(32)"`
	Phone      string    `json:"phone" xorm:"comment('联系电话') varchar(18)"`
	Remark     string    `json:"remark" xorm:"comment('备注信息') varchar(255)"`
	StartTime  time.Time `json:"start_time" xorm:"comment('起租时间') datetime"`
	EndTime    time.Time `json:"end_time" xorm:"comment('结束时间') datetime"`
	Email      string    `json:"email" xorm:"comment('安全邮箱') varchar(255)"`
}

//新增页面请求参数
type TenantAddReq struct {
	DelFlag   string `form:"delFlag" `
	Name      string `form:"name" binding:"required"`
	Address   string `form:"address" `
	Manager   string `form:"manager" `
	Phone     string `form:"phone" `
	Remark    string `form:"remark" `
	StartTime string `form:"startTime" `
	EndTime   string `form:"endTime" `
	Email     string `form:"email" `
}

//修改页面请求参数
type TenantEditReq struct {
	Id        int64  `form:"id" binding:"required"`
	Name      string `form:"name" binding:"required"`
	Address   string `form:"address" `
	Manager   string `form:"manager" `
	Phone     string `form:"phone" `
	Remark    string `form:"remark" `
	StartTime string `form:"startTime" `
	EndTime   string `form:"endTime" `
	Email     string `form:"email" `
}

//分页请求参数
type TenantSelectPageReq struct {
	Name      string `form:"name"`      //商户名称
	Address   string `form:"address"`   //联系地址
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
}
