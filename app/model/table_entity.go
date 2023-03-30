package model

import (
	"time"
)

type TableEntity struct {
	TableId        int64     `json:"table_id" xorm:"not null pk autoincr comment('编号') BIGINT(20)"`
	TableName      string    `json:"table_name" xorm:"default '' comment('表名称') VARCHAR(200)"`
	TableComment   string    `json:"table_comment" xorm:"default '' comment('表描述') VARCHAR(500)"`
	ClassName      string    `json:"class_name" xorm:"default '' comment('实体类名称') VARCHAR(100)"`
	TplCategory    string    `json:"tpl_category" xorm:"default 'crud' comment('使用的模板（crud单表操作 tree树表操作）') VARCHAR(200)"`
	PackageName    string    `json:"package_name" xorm:"comment('生成包路径') VARCHAR(100)"`
	ModuleName     string    `json:"module_name" xorm:"comment('生成模块名') VARCHAR(30)"`
	BusinessName   string    `json:"business_name" xorm:"comment('生成业务名') VARCHAR(30)"`
	FunctionName   string    `json:"function_name" xorm:"comment('生成功能名') VARCHAR(50)"`
	FunctionAuthor string    `json:"function_author" xorm:"comment('生成功能作者') VARCHAR(50)"`
	Options        string    `json:"options" xorm:"comment('其它生成选项') VARCHAR(1000)"`
	CreateBy       string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy       string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// Fill with you ideas below.

// Entity is the golang structure for table gen_table.
type TableEntityExtend struct {
	TableEntity    `xorm:"extends"`
	TreeCode       string              `xorm:"-"` // 树编码字段
	TreeParentCode string              `xorm:"-"` // 树父编码字段
	TreeName       string              `xorm:"-"` // 树名称字段
	Columns        []TableColumnEntity `xorm:"-"` // 表列信息
	PkColumn       TableColumnEntity   `xorm:"-"` // 表列信息
}

type TableParams struct {
	TreeCode       string `form:"treeCode"`
	TreeParentCode string `form:"treeParentCode"`
	TreeName       string `form:"treeName"`
}

//修改页面请求参数
type TableEditReq struct {
	TableId        int64  `form:"tableId" binding:"required"`
	TableName      string `form:"tableName"  binding:"required"`
	TableComment   string `form:"tableComment"  binding:"required"`
	ClassName      string `form:"className" binding:"required"`
	FunctionAuthor string `form:"functionAuthor"  binding:"required"`
	TplCategory    string `form:"tplCategory"`
	PackageName    string `form:"packageName" binding:"required"`
	ModuleName     string `form:"moduleName" binding:"required"`
	BusinessName   string `form:"businessName" binding:"required"`
	FunctionName   string `form:"functionName" binding:"required"`
	Remark         string `form:"remark"`
	Params         string `form:"params"`
	Columns        string `form:"columns"`
}

//分页请求参数
type TableSelectPageReq struct {
	TableName    string `form:"tableName"`    //表名称
	TableComment string `form:"tableComment"` //表描述
	BeginTime    string `form:"beginTime"`    //开始时间
	EndTime      string `form:"endTime"`      //结束时间
	PageNum      int    `form:"pageNum"`      //当前页码
	PageSize     int    `form:"pageSize"`     //每页数
}
