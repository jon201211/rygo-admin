/* ==========================================================================
 RYGO自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
 生成日期：2020-03-27 04:35:17 +0800 CST
 ==========================================================================*/

package config

import (
	"yj-app/app/db"
	"time"
)

// 数据表映射结构体
type ConfigEntity struct { 
	 ConfigId       int         `json:"config_id" xorm:"not null pk autoincr comment('参数主键11') int(5)"`    
	 ConfigName    string         `json:"config_name" xorm:"comment('参数名称111') varchar(100)"`    
	 ConfigKey    string         `json:"config_key" xorm:"comment('参数键名111') varchar(100)"`    
	 ConfigValue    string         `json:"config_value" xorm:"comment('参数键值') varchar(500)"`    
	 ConfigType    string         `json:"config_type" xorm:"comment('系统内置（Y是 N否）') char(1)"`    
	 CreateBy    string         `json:"create_by" xorm:"comment('创建者') varchar(64)"`    
	 CreateTime    time.Time         `json:"create_time" xorm:"comment('创建时间') datetime"`    
	 UpdateBy    string         `json:"update_by" xorm:"comment('更新者') varchar(64)"`    
	 UpdateTime    time.Time         `json:"update_time" xorm:"comment('更新时间') datetime"`    
	 Remark    string         `json:"remark" xorm:"comment('备注') varchar(500)"`    
}

//新增页面请求参数
type ConfigAddReq struct { 
	 
	 ConfigName  string   `form:"configName" binding:"required"`  
	 ConfigKey  string   `form:"configKey" `  
	 ConfigValue  string   `form:"configValue" `  
	 ConfigType  string   `form:"configType" `  
	 
	 
	 
	 
	 Remark  string   `form:"remark" `  
}

//修改页面请求参数
type ConfigEditReq struct {
      ConfigId    int  `form:"configId" binding:"required"`    
      ConfigName  string `form:"configName" binding:"required"`   
      ConfigKey  string `form:"configKey" `   
      ConfigValue  string `form:"configValue" `   
      ConfigType  string `form:"configType" `           
      Remark  string `form:"remark" `  
}

//分页请求参数 
type ConfigSelectPageReq struct {  
	ConfigId  int `form:"configId"` //参数主键11   
	ConfigName  string `form:"configName"` //参数名称111   
	ConfigKey  string `form:"configKey"` //参数键名111   
	ConfigValue  string `form:"configValue"` //参数键值   
	ConfigType  string `form:"configType"` //系统内置（Y是 N否）            
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}

