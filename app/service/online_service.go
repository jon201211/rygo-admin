package service

import (
	"rygo/app/dao"
	"rygo/app/model"
	"rygo/app/utils/page"
	"strings"
)

var OnlineService = newOnlineService()

func newOnlineService() *onlineService {
	return &onlineService{}
}

type onlineService struct {
}

func (s *onlineService) Insert(r *model.UserOnline) (int64, error) {
	return dao.UserOnlineDao.Insert(r)
}
func (s *onlineService) Delete(r *model.UserOnline) (int64, error) {
	return dao.UserOnlineDao.Insert(r)
}

//根据主键查询数据
func (s *onlineService) SelectRecordById(id string) (*model.UserOnline, error) {
	entity := &model.UserOnline{Sessionid: id}
	_, err := dao.UserOnlineDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *onlineService) DeleteRecordById(id string) bool {
	entity := &model.UserOnline{Sessionid: id}
	result, err := dao.UserOnlineDao.Delete(entity)
	if err == nil && result > 0 {
		return true
	}

	return false
}

//批量删除数据记录
func (s *onlineService) DeleteRecordByIds(ids string) int64 {
	idarr := strings.Split(ids, ",")
	result, _ := dao.UserOnlineDao.DeleteBatch(idarr...)
	return result
}

//批量删除数据
func (s *onlineService) DeleteRecordNotInIds(ids []string) int64 {
	result, _ := dao.UserOnlineDao.DeleteNotIn(ids...)
	return result
}

//添加数据
func (s *onlineService) AddSave(entity model.UserOnline) (int64, error) {
	return dao.UserOnlineDao.Insert(&entity)
}

//根据条件查询数据
func (s *onlineService) SelectListAll(params *model.UserOnlineSelectPageReq) ([]model.UserOnline, error) {
	return dao.UserOnlineDao.SelectListAll(params)
}

//根据条件分页查询数据
func (s *onlineService) SelectListByPage(params *model.UserOnlineSelectPageReq) ([]model.UserOnline, *page.Paging, error) {
	return dao.UserOnlineDao.SelectListByPage(params)
}
