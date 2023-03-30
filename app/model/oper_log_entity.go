package model

import (
	"time"
)

type OperLogEntity struct {
	OperId        int64     `json:"oper_id" xorm:"not null pk autoincr comment('日志主键') BIGINT(20)"`
	Title         string    `json:"title" xorm:"default '' comment('模块标题') VARCHAR(50)"`
	BusinessType  int       `json:"business_type" xorm:"default 0 comment('业务类型（0其它 1新增 2修改 3删除）') INT(2)"`
	Method        string    `json:"method" xorm:"default '' comment('方法名称') VARCHAR(100)"`
	RequestMethod string    `json:"request_method" xorm:"default '' comment('请求方式') VARCHAR(10)"`
	OperatorType  int       `json:"operator_type" xorm:"default 0 comment('操作类别（0其它 1后台用户 2手机端用户）') INT(1)"`
	OperName      string    `json:"oper_name" xorm:"default '' comment('操作人员') VARCHAR(50)"`
	DeptName      string    `json:"dept_name" xorm:"default '' comment('部门名称') VARCHAR(50)"`
	OperUrl       string    `json:"oper_url" xorm:"default '' comment('请求URL') VARCHAR(255)"`
	OperIp        string    `json:"oper_ip" xorm:"default '' comment('主机地址') VARCHAR(50)"`
	OperLocation  string    `json:"oper_location" xorm:"default '' comment('操作地点') VARCHAR(255)"`
	OperParam     string    `json:"oper_param" xorm:"default '' comment('请求参数') VARCHAR(2000)"`
	JsonResult    string    `json:"json_result" xorm:"default '' comment('返回参数') VARCHAR(2000)"`
	Status        int       `json:"status" xorm:"default 0 comment('操作状态（0正常 1异常）') INT(1)"`
	ErrorMsg      string    `json:"error_msg" xorm:"default '' comment('错误消息') VARCHAR(2000)"`
	OperTime      time.Time `json:"oper_time" xorm:"comment('操作时间') DATETIME"`
}

//
//查询列表请求参数
type OperLogSelectPageReq struct {
	Title         string `form:"title"`         //系统模块
	OperName      string `form:"operName"`      //操作人员
	BusinessTypes int    `form:"businessTypes"` //操作类型
	Status        string `form:"status"`        //操作类型
	BeginTime     string `form:"beginTime"`     //数据范围
	EndTime       string `form:"endTime"`       //开始时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}
