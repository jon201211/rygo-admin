package model

import (
	"time"
)

type SysDept struct {
	DeptId     int64     `json:"dept_id" xorm:"not null pk autoincr comment('部门id') BIGINT(20)"`
	ParentId   int64     `json:"parent_id" xorm:"default 0 comment('父部门id') BIGINT(20)"`
	Ancestors  string    `json:"ancestors" xorm:"default '' comment('祖级列表') VARCHAR(50)"`
	DeptName   string    `json:"dept_name" xorm:"default '' comment('部门名称') VARCHAR(30)"`
	OrderNum   int       `json:"order_num" xorm:"default 0 comment('显示顺序') INT(4)"`
	Leader     string    `json:"leader" xorm:"comment('负责人') VARCHAR(20)"`
	Phone      string    `json:"phone" xorm:"comment('联系电话') VARCHAR(11)"`
	Email      string    `json:"email" xorm:"comment('邮箱') VARCHAR(50)"`
	Status     string    `json:"status" xorm:"default '0' comment('部门状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `json:"del_flag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`

	TenantId int64 `json:"tenant_id" xorm:"default 0 comment('租户id') BIGINT(20)"`
}

// Fill with you ideas below.

// Entity is the golang structure for table sys_dept.
type SysDeptExtend struct {
	SysDept    `xorm:"extends"`
	ParentName string `json:"parentName"`
}

//分页请求参数
type DeptSelectPageReq struct {
	ParentId  int64  `form:"parentId"`      //父部门ID
	DeptName  string `form:"deptName"`      //部门名称
	Status    string `form:"status"`        //状态
	PageNum   int    `form:"pageNum"`       //当前页码
	PageSize  int    `form:"pageSize"`      //每页数
	SortName  string `form:"orderByColumn"` //排序字段
	SortOrder string `form:"isAsc"`         //排序方式
	BeginTime string `form:"beginTime"`     //开始时间
	EndTime   string `form:"endTime"`       //结束时间
	TenantId  int64  `form:"tenantId"`      //结束时间
}

//新增页面请求参数
type DeptAddReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Leader   string `form:"leader"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   string `form:"status"`
	TenantId int64  `form:"tenantId"` //结束时间
}

//修改页面请求参数
type DeptEditReq struct {
	DeptId int64 `form:"deptId" binding:"required"`
	DeptAddReq
}

//检查菜单名称请求参数
type CheckDeptNameReq struct {
	DeptId   int64  `form:"deptId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}

//检查菜单名称请求参数
type CheckDeptNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}
