package service

import (
	"errors"

	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/page"
	"time"

	"github.com/gin-gonic/gin"
)

var PostService = newPostService()

func newPostService() *postService {
	return &postService{}
}

type postService struct {
}

//根据主键查询数据
func (s *postService) SelectRecordById(id int64) (*model.SysPost, error) {
	entity := &model.SysPost{PostId: id}
	_, err := dao.PostDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *postService) DeleteRecordById(id int64) bool {
	entity := &model.SysPost{PostId: id}
	rs, err := dao.PostDao.Delete(entity)
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (s *postService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dao.PostDao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (s *postService) AddSave(req *model.PostAddReq, ctx *gin.Context) (int64, error) {
	var entity model.SysPost
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.PostSort = req.PostSort
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.PostDao.Insert(&entity)
	return entity.PostId, err
}

//修改数据
func (s *postService) EditSave(req *model.PostEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.SysPost{PostId: req.PostId}
	ok, err := dao.PostDao.FindOne(entity)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.PostSort = req.PostSort
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return dao.PostDao.Update(entity)
}

//根据条件分页查询角色数据
func (s *postService) SelectListAll(params *model.PostSelectPageReq) ([]model.PostEntityFlag, error) {
	return dao.PostDao.SelectListAll(params)
}

//根据条件分页查询角色数据
func (s *postService) SelectListByPage(params *model.PostSelectPageReq) ([]model.SysPost, *page.Paging, error) {
	return dao.PostDao.SelectListByPage(params)
}

// 导出excel
func (s *postService) Export(param *model.PostSelectPageReq) (string, error) {
	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	col := []string{"post_id", "post_name", "post_code", "post_sort", "stat"}
	return dao.PostDao.SelectListExport(param, head, col)
}

//根据用户ID查询岗位
func (s *postService) SelectPostsByUserId(userId int64) ([]model.PostEntityFlag, error) {
	var paramsPost *model.PostSelectPageReq
	postAll, err := dao.PostDao.SelectListAll(paramsPost)

	if err != nil || postAll == nil {
		return nil, errors.New("未查询到岗位数据")
	}

	userPost, err := dao.PostDao.SelectPostsByUserId(userId)

	if err != nil || userPost == nil {
		return nil, errors.New("未查询到用户岗位数据")
	} else {
		for i := range postAll {
			for j := range userPost {
				if userPost[j].PostId == postAll[i].PostId {
					postAll[i].Flag = true
					break
				}
			}
		}
	}

	return postAll, nil
}

//检查角色名是否唯一
func (s *postService) CheckPostNameUniqueAll(postName string) string {
	post, err := dao.PostDao.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位名称是否唯一
func (s *postService) CheckPostNameUnique(postName string, postId int64) string {
	post, err := dao.PostDao.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func (s *postService) CheckPostCodeUniqueAll(postCode string) string {
	post, err := dao.PostDao.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func (s *postService) CheckPostCodeUnique(postCode string, postId int64) string {
	post, err := dao.PostDao.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 && post.PostId != postId {
		return "1"
	}
	return "0"
}
