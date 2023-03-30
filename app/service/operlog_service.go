package service

import (
	"encoding/json"
	"errors"
	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/ip"
	"rygo/app/utils/page"
	"time"

	"github.com/gin-gonic/gin"
)

var OperlogService = newOperlogService()

func newOperlogService() *operlogService {
	return &operlogService{}
}

type operlogService struct {
}

//新增记录
func (s *operlogService) Add(ctx *gin.Context, title, inContent string, outContent *model.CommonRes) error {
	user := UserService.GetProfile(ctx)
	if user == nil {
		return errors.New("用户未登陆")
	}

	var operLog model.OperLogEntity

	outJson, _ := json.Marshal(outContent)
	outJsonStr := string(outJson)

	operLog.Title = title
	operLog.OperParam = inContent
	operLog.JsonResult = outJsonStr
	operLog.BusinessType = int(outContent.Btype)
	//操作类别（0其它 1后台用户 2手机端用户）
	operLog.OperatorType = 1
	//操作状态（0正常 1异常）
	if outContent.Code == 0 {
		operLog.Status = 0
	} else {
		operLog.Status = 1
	}

	operLog.OperName = user.LoginName
	operLog.RequestMethod = ctx.Request.Method

	//获取用户部门
	dept := DeptService.SelectDeptById(user.DeptId)

	if dept != nil {
		operLog.DeptName = dept.DeptName
	} else {
		operLog.DeptName = ""
	}

	operLog.OperUrl = ctx.Request.URL.Path
	operLog.Method = ctx.Request.Method
	operLog.OperIp = ctx.ClientIP()

	operLog.OperLocation = ip.GetCityByIp(operLog.OperIp)
	operLog.OperTime = time.Now()

	_, err := dao.OperLogDao.Insert(&operLog)
	return err
}

// 根据条件分页查询用户列表
func (s *operlogService) SelectPageList(param *model.OperLogSelectPageReq) (*[]model.OperLogEntity, *page.Paging, error) {
	return dao.OperLogDao.SelectPageList(param)
}

//根据主键查询用户信息
func (s *operlogService) SelectRecordById(id int64) (*model.OperLogEntity, error) {
	entity := &model.OperLogEntity{OperId: id}
	_, err := dao.OperLogDao.FindOne(entity)
	return entity, err
}

//根据主键删除用户信息
func (s *operlogService) DeleteRecordById(id int64) bool {
	entity := &model.OperLogEntity{OperId: id}
	result, err := dao.OperLogDao.Delete(entity)
	if err == nil && result > 0 {
		return true
	}

	return false
}

//批量删除记录
func (s *operlogService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, _ := dao.OperLogDao.DeleteBatch(idarr...)
	return result
}

//清空记录
func (s *operlogService) DeleteRecordAll() (int64, error) {
	return dao.OperLogDao.DeleteAll()
}

// 导出excel
func (s *operlogService) Export(param *model.OperLogSelectPageReq) (string, error) {
	head := []string{"日志主键", "模块标题", "业务类型", "方法名称", "请求方式", "操作类别", "操作人员", "部门名称", "请求URL", "主机地址", "操作地点", "请求参数", "返回参数", "操作状态", "操作时间"}
	col := []string{"oper_id", "title", "business_type", "method", "request_method", "operator_type", "oper_name", "dept_name", "oper_url", "oper_ip", "oper_location", "oper_param", "json_result", "status", "error_msg", "oper_time"}
	return dao.OperLogDao.SelectExportList(param, head, col)
}
