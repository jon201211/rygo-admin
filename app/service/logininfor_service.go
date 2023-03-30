package service

import (
	"rygo/app/cache"
	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/gconv"
	"rygo/app/utils/page"
	"time"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

var LogininforService = newLogininforService()

func newLogininforService() *logininforService {
	return &logininforService{}
}

type logininforService struct {
}

// 根据条件分页查询用户列表
func (s *logininforService) SelectPageList(param *model.LogininforSelectPageReq) (*[]model.LogininforEntity, *page.Paging, error) {
	return dao.LogininforDao.SelectPageList(param)
}

//根据主键查询用户信息
func (s *logininforService) SelectRecordById(id int64) (*model.LogininforEntity, error) {
	entity := &model.LogininforEntity{InfoId: id}
	_, err := dao.LogininforDao.FindOne(entity)
	return entity, err
}

//根据主键删除用户信息
func (s *logininforService) DeleteRecordById(id int64) bool {
	entity := &model.LogininforEntity{InfoId: id}
	result, err := dao.LogininforDao.Delete(entity)
	if err == nil && result > 0 {
		return true
	}

	return false
}

//批量删除记录
func (s *logininforService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, _ := dao.LogininforDao.DeleteBatch(idarr...)
	return result
}

//清空记录
func (s *logininforService) DeleteRecordAll() (int64, error) {
	return dao.LogininforDao.DeleteAll()
}

func (s *logininforService) Insert(r *model.LogininforEntity) (int64, error) {
	return dao.LogininforDao.Insert(r)
}

// 导出excel
func (s *logininforService) Export(param *model.LogininforSelectPageReq) (string, error) {
	head := []string{"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"}
	col := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	return dao.LogininforDao.SelectExportList(param, head, col)
}

//记录密码尝试次数
func (s *logininforService) SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := cache.Instance().Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	curTimes = curTimes + 1
	cache.Instance().Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= 5 {
		s.Lock(loginName)
	}
	return curTimes
}

//记录密码尝试次数
func (s *logininforService) GetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := cache.Instance().Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	return curTimes
}

//移除密码错误次数
func (s *logininforService) RemovePasswordCounts(loginName string) {
	cache.Instance().Delete(USER_NOPASS_TIME + loginName)
}

//锁定账号
func (s *logininforService) Lock(loginName string) {
	cache.Instance().Set(USER_LOCK+loginName, true, 30*time.Minute)
}

//解除锁定
func (s *logininforService) Unlock(loginName string) {
	cache.Instance().Delete(USER_LOCK + loginName)
}

//检查账号是否锁定
func (s *logininforService) CheckLock(loginName string) bool {
	result := false
	rs, _ := cache.Instance().Get(USER_LOCK + loginName)
	if rs != nil {
		result = true
	}
	return result
}
