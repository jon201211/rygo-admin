package model

import (
	"time"
)

type SysDictData struct {
	DictCode   int64     `json:"dict_code" xorm:"not null pk autoincr comment('字典编码') BIGINT(20)"`
	DictSort   int       `json:"dict_sort" xorm:"default 0 comment('字典排序') INT(4)"`
	DictLabel  string    `json:"dict_label" xorm:"default '' comment('字典标签') VARCHAR(100)"`
	DictValue  string    `json:"dict_value" xorm:"default '' comment('字典键值') VARCHAR(100)"`
	DictType   string    `json:"dict_type" xorm:"default '' comment('字典类型') VARCHAR(100)"`
	CssClass   string    `json:"css_class" xorm:"comment('样式属性（其他样式扩展）') VARCHAR(100)"`
	ListClass  string    `json:"list_class" xorm:"comment('表格回显样式') VARCHAR(100)"`
	IsDefault  string    `json:"is_default" xorm:"default 'N' comment('是否默认（Y是 N否）') CHAR(1)"`
	Status     string    `json:"status" xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// Fill with you ideas below.
//新增页面请求参数
type DictDataAddReq struct {
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"  binding:"required"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

//修改页面请求参数
type DictDataEditReq struct {
	DictCode  int64  `form:"dictCode" binding:"required"`
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

//分页请求参数
type DictDataSelectPageReq struct {
	DictType  string `form:"dictType"`  //字典名称
	DictLabel string `form:"dictLabel"` //字典标签
	Status    string `form:"status"`    //状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
}
