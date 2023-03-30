package model

import (
	"time"
)

type SysDictType struct {
	DictId     int64     `json:"dict_id" xorm:"not null pk autoincr comment('字典主键') BIGINT(20)"`
	DictName   string    `json:"dict_name" xorm:"default '' comment('字典名称') VARCHAR(100)"`
	DictType   string    `json:"dict_type" xorm:"default '' comment('字典类型') unique VARCHAR(100)"`
	Status     string    `json:"status" xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// Fill with you ideas below.
//新增页面请求参数
type DictTypeAddReq struct {
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

//修改页面请求参数
type DictTypeEditReq struct {
	DictId   int64  `form:"dictId" binding:"required"`
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

//分页请求参数
type DictTypeSelectPageReq struct {
	DictName      string `form:"dictName"`      //字典名称
	DictType      string `form:"dictType"`      //字典类型
	Status        string `form:"status"`        //字典状态
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
}

//检查字典类型请求参数
type CheckDictTypeReq struct {
	DictId   int64  `form:"dictId"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
}

//检查字典类型请求参数
type CheckDictTypeALLReq struct {
	DictType string `form:"dictType"  binding:"required"`
}
