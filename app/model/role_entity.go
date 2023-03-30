package model

import (
	"time"
)

type RoleEntity struct {
	RoleId     int64     `json:"role_id" xorm:"not null pk autoincr comment('角色ID') BIGINT(20)"`
	RoleName   string    `json:"role_name" xorm:"not null comment('角色名称') VARCHAR(30)"`
	RoleKey    string    `json:"role_key" xorm:"not null comment('角色权限字符串') VARCHAR(100)"`
	RoleSort   int       `json:"role_sort" xorm:"not null comment('显示顺序') INT(4)"`
	DataScope  string    `json:"data_scope" xorm:"default '1' comment('数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）') CHAR(1)"`
	Status     string    `json:"status" xorm:"not null comment('角色状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `json:"del_flag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// Entity is the golang structure for table sys_role.
type RoleEntityFlag struct {
	RoleId     int64     `json:"role_id" xorm:"not null pk autoincr comment('角色ID') BIGINT(20)"`
	RoleName   string    `json:"role_name" xorm:"not null comment('角色名称') VARCHAR(30)"`
	RoleKey    string    `json:"role_key" xorm:"not null comment('角色权限字符串') VARCHAR(100)"`
	RoleSort   int       `json:"role_sort" xorm:"not null comment('显示顺序') INT(4)"`
	DataScope  string    `json:"data_scope" xorm:"default '1' comment('数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）') CHAR(1)"`
	Status     string    `json:"status" xorm:"not null comment('角色状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `json:"del_flag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	Flag       bool      `json:"flag" xorm:"comment('标记') BOOL"`
}

//数据权限保存请求参数
type RoleDataScopeReq struct {
	RoleId    int64  `form:"roleId"  binding:"required"`
	RoleName  string `form:"roleName"  binding:"required"`
	RoleKey   string `form:"roleKey"  binding:"required"`
	DataScope string `form:"dataScope"  binding:"required"`
	DeptIds   string `form:"deptIds"`
}

//检查角色名称请求参数
type RoleCheckRoleNameReq struct {
	RoleId   int64  `form:"roleId"  binding:"required"`
	RoleName string `form:"roleName"  binding:"required"`
}

//检查权限字符请求参数
type RoleCheckRoleKeyReq struct {
	RoleId  int64  `form:"roleId"  binding:"required"`
	RoleKey string `form:"roleKey"  binding:"required"`
}

//检查角色名称请求参数
type RoleCheckRoleNameALLReq struct {
	RoleName string `form:"roleName"  binding:"required"`
}

//检查权限字符请求参数
type RoleCheckRoleKeyALLReq struct {
	RoleKey string `form:"roleKey"  binding:"required"`
}

//分页请求参数
type RoleSelectPageReq struct {
	RoleName      string `form:"roleName"`      //角色名称
	Status        string `form:"status"`        //状态
	RoleKey       string `form:"roleKey"`       //角色键
	DataScope     string `form:"dataScope"`     //数据范围
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}

//新增页面请求参数
type RoleAddReq struct {
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort string `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds""`
}

//修改页面请求参数
type RoleEditReq struct {
	RoleId   int64  `form:"roleId" binding:"required"`
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort string `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds"`
}
