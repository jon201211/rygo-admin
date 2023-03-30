package model

import (
	"time"
)

type MenuEntity struct {
	MenuId     int64     `json:"menu_id" xorm:"not null pk autoincr comment('菜单ID') BIGINT(20)"`
	MenuName   string    `json:"menu_name" xorm:"not null comment('菜单名称') VARCHAR(50)"`
	ParentId   int64     `json:"parent_id" xorm:"default 0 comment('父菜单ID') BIGINT(20)"`
	OrderNum   int       `json:"order_num" xorm:"default 0 comment('显示顺序') INT(4)"`
	Url        string    `json:"url" xorm:"default '#' comment('请求地址') VARCHAR(200)"`
	Target     string    `json:"target" xorm:"default '' comment('打开方式（menuItem页签 menuBlank新窗口）') VARCHAR(20)"`
	MenuType   string    `json:"menu_type" xorm:"default '' comment('菜单类型（M目录 C菜单 F按钮）') CHAR(1)"`
	Visible    string    `json:"visible" xorm:"default '0' comment('菜单状态（0显示 1隐藏）') CHAR(1)"`
	Perms      string    `json:"perms" xorm:"comment('权限标识') VARCHAR(100)"`
	Icon       string    `json:"icon" xorm:"default '#' comment('菜单图标') VARCHAR(100)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
}

// Entity is the golang structure for table sys_menu.
type MenuEntityExtend struct {
	MenuEntity `xorm:"extends"`
	ParentName string             `json:"parentName"` // 父菜单名称
	Children   []MenuEntityExtend `json:"children"`   // 子菜单
}

//检查菜单名称请求参数
type MenuCheckMenuNameReq struct {
	MenuId   int64  `form:"menuId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

//检查菜单名称请求参数
type MenuCheckMenuNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

//分页请求参数
type MenuSelectPageReq struct {
	MenuName  string `form:"menuName"`      //菜单名称
	Visible   string `form:"visible"`       //状态
	BeginTime string `form:"beginTime"`     //开始时间
	EndTime   string `form:"endTime"`       //结束时间
	PageNum   int    `form:"pageNum"`       //当前页码
	PageSize  int    `form:"pageSize"`      //每页数
	SortName  string `form:"orderByColumn"` //排序字段
	SortOrder string `form:"isAsc"`         //排序方式
}

//新增页面请求参数
type MenuAddReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}

//修改页面请求参数
type MenuEditReq struct {
	MenuId   int64  `form:"menuId" binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}
