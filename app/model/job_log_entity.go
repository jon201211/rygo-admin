package model

import (
	"time"

	_ "github.com/go-xorm/xorm"
)

type JobLogEntity struct {
	JobLogId      int64     `json:"job_log_id" xorm:"not null pk autoincr comment('任务日志ID') BIGINT(20)"`
	JobName       string    `json:"job_name" xorm:"not null comment('任务名称') VARCHAR(64)"`
	JobGroup      string    `json:"job_group" xorm:"not null comment('任务组名') VARCHAR(64)"`
	InvokeTarget  string    `json:"invoke_target" xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	JobMessage    string    `json:"job_message" xorm:"comment('日志信息') VARCHAR(500)"`
	Status        string    `json:"status" xorm:"default '0' comment('执行状态（0正常 1失败）') CHAR(1)"`
	ExceptionInfo string    `json:"exception_info" xorm:"default '' comment('异常信息') VARCHAR(2000)"`
	CreateTime    time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
}

//分页请求参数
type JobLogSelectPageReq struct {
	JobLogId      int64  `form:"jobLogId"`      //任务日志ID
	JobName       string `form:"jobName"`       //任务名称
	JobGroup      string `form:"jobGroup"`      //任务组名
	InvokeTarget  string `form:"invokeTarget"`  //调用目标字符串
	JobMessage    string `form:"jobMessage"`    //日志信息
	Status        string `form:"status"`        //执行状态（0正常 1失败）
	ExceptionInfo string `form:"exceptionInfo"` //异常信息
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	PageNum       int    `form:"pageNum"`       //当前页码
	PageSize      int    `form:"pageSize"`      //每页数
}
