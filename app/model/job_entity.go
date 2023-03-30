package model

import (
	"time"
)

type SysJob struct {
	JobId          int64     `json:"job_id" xorm:"not null pk autoincr comment('任务ID') BIGINT(20)"`
	JobName        string    `json:"job_name" xorm:"not null default '' comment('任务名称') VARCHAR(64)"`
	JobParams      string    `json:"job_params" xorm:"default '' comment('参数') VARCHAR(200)"`
	JobGroup       string    `json:"job_group" xorm:"not null default 'DEFAULT' comment('任务组名') VARCHAR(64)"`
	InvokeTarget   string    `json:"invoke_target" xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	CronExpression string    `json:"cron_expression" xorm:"default '' comment('cron执行表达式') VARCHAR(255)"`
	MisfirePolicy  string    `json:"misfire_policy" xorm:"default '1' comment('计划执行策略（1多次执行 2执行一次）') VARCHAR(20)"`
	Concurrent     string    `json:"concurrent" xorm:"default '1' comment('是否并发执行（0允许 1禁止）') CHAR(1)"`
	Status         string    `json:"status" xorm:"default '0' comment('状态（0正常 1暂停）') CHAR(1)"`
	CreateBy       string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy       string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark         string    `json:"remark" xorm:"default '' comment('备注信息') VARCHAR(500)"`
}

//新增页面请求参数
type JobAddReq struct {
	JobName        string `form:"jobName" `
	JobParams      string `form:"jobParams"` // 任务参数
	JobGroup       string `form:"jobGroup" `
	InvokeTarget   string `form:"invokeTarget" `
	CronExpression string `form:"cronExpression" `
	MisfirePolicy  string `form:"misfirePolicy" `
	Concurrent     string `form:"concurrent" `
	Status         string `form:"status" binding:"required"`
	Remark         string `form:"remark" `
}

//修改页面请求参数
type JobEditReq struct {
	JobName        string `form:"jobName" `
	JobParams      string `form:"jobParams"` // 任务参数
	JobGroup       string `form:"jobGroup" `
	JobId          int64  `form:"jobId" binding:"required"`
	InvokeTarget   string `form:"invokeTarget" `
	CronExpression string `form:"cronExpression" `
	MisfirePolicy  string `form:"misfirePolicy" `
	Concurrent     string `form:"concurrent" `
	Status         string `form:"status" binding:"required"`
	Remark         string `form:"remark" `
}

//分页请求参数
type JobSelectPageReq struct {
	JobId          int64  `form:"jobId"`          //任务ID
	JobName        string `form:"jobName"`        //任务名称
	JobGroup       string `form:"jobGroup"`       //任务组名
	InvokeTarget   string `form:"invokeTarget"`   //调用目标字符串
	CronExpression string `form:"cronExpression"` //cron执行表达式
	MisfirePolicy  string `form:"misfirePolicy"`  //计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string `form:"concurrent"`     //是否并发执行（0允许 1禁止）
	Status         string `form:"status"`         //状态（0正常 1暂停）
	BeginTime      string `form:"beginTime"`      //开始时间
	EndTime        string `form:"endTime"`        //结束时间
	PageNum        int    `form:"pageNum"`        //当前页码
	PageSize       int    `form:"pageSize"`       //每页数
	OrderByColumn  string `form:"orderByColumn"`  //排序字段
	IsAsc          string `form:"isAsc"`          //排序方式
}
