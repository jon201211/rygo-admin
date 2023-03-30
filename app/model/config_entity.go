package model

import (
	_ "rygo/app/db"
	"time"
)

/*
工具自动生成方法:使用xorm工具
安装工具:go get github.com/go-xorm/cmd/xorm
命令行输入:xorm reverse mysql root:密码@/xorm?charset=utf8 /home/tym/golib/src/github.com/go-xorm/cmd/xorm/templates/goxorm/
/home/tym/golib/src/github.com/go-xorm/cmd/xorm/templates/goxorm/这一串是模版的位置(GOPATH/src里)
不写生成路径会在你的目录下建一个model,对应文件生成在model中
*/

type SysConfig struct {
	ConfigId    int64     `json:"config_id" xorm:"not null pk autoincr comment('参数主键') INT(5)"`
	ConfigName  string    `json:"config_name" xorm:"default '' comment('参数名称') VARCHAR(100)"`
	ConfigKey   string    `json:"config_key" xorm:"default '' comment('参数键名') VARCHAR(100)"`
	ConfigValue string    `json:"config_value" xorm:"default '' comment('参数键值') VARCHAR(500)"`
	ConfigType  string    `json:"config_type" xorm:"default 'N' comment('系统内置（Y是 N否）') CHAR(1)"`
	CreateBy    string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// Fill with you ideas below.
//新增页面请求参数
type ConfigAddReq struct {
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

//修改页面请求参数
type ConfigEditReq struct {
	ConfigId    int64  `form:"configId" binding:"required"`
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

//分页请求参数
type ConfigSelectPageReq struct {
	ConfigName string `form:"configName"` //参数名称
	ConfigKey  string `form:"configKey"`  //参数键名
	ConfigType string `form:"configType"` //状态
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}

//检查参数键名请求参数
type CheckConfigKeyReq struct {
	ConfigId  int64  `form:"configId"  binding:"required"`
	ConfigKey string `form:"configKey"  binding:"required"`
}

//检查参数键名请求参数
type CheckPostCodeALLReq struct {
	ConfigKey string `form:"configKey"  binding:"required"`
}
