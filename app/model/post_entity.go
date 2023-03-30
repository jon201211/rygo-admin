package model

import (
	"time"
)

type SysPost struct {
	PostId     int64     `json:"post_id" xorm:"not null pk autoincr comment('岗位ID') BIGINT(20)"`
	PostCode   string    `json:"post_code" xorm:"not null comment('岗位编码') VARCHAR(64)"`
	PostName   string    `json:"post_name" xorm:"not null comment('岗位名称') VARCHAR(50)"`
	PostSort   int       `json:"post_sort" xorm:"not null comment('显示顺序') INT(4)"`
	Status     string    `json:"status" xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`

	TenantId int64 `json:"tenant_id" xorm:"default 0 comment('租户id') BIGINT(20)"`
}

// Fill with you ideas below.
// SysPost is the golang structure for table sys_post.
type PostEntityFlag struct {
	PostId     int64     `json:"post_id" xorm:"not null pk autoincr comment('岗位ID') BIGINT(20)"`
	PostCode   string    `json:"post_code" xorm:"not null comment('岗位编码') VARCHAR(64)"`
	PostName   string    `json:"post_name" xorm:"not null comment('岗位名称') VARCHAR(50)"`
	PostSort   int       `json:"post_sort" xorm:"not null comment('显示顺序') INT(4)"`
	Status     string    `json:"status" xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	Flag       bool      `json:"flag"` // 标记
}

//新增页面请求参数
type PostAddReq struct {
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

//修改页面请求参数
type PostEditReq struct {
	PostId   int64  `form:"postId" binding:"required"`
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

//分页请求参数
type PostSelectPageReq struct {
	PostCode      string `form:"postCode"`      //岗位编码
	Status        string `form:"status"`        //状态
	PostName      string `form:"postName"`      //岗位名称
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
}

//检查编码请求参数
type PostCheckPostCodeReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
}

//检查编码请求参数
type PostCheckPostCodeALLReq struct {
	PostCode string `form:"postCode"  binding:"required"`
}

//检查名称请求参数
type PostCheckPostNameReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostName string `form:"postName"  binding:"required"`
}

//检查名称请求参数
type PostCheckPostNameALLReq struct {
	PostName string `form:"postName"  binding:"required"`
}
