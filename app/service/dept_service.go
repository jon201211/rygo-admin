package service

import (
	"errors"
	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/gconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var DeptService = newDeptService()

func newDeptService() *deptService {
	return &deptService{}
}

type deptService struct {
}

//新增保存信息
func (s *deptService) AddSave(req *model.DeptAddReq, ctx *gin.Context) (int64, error) {
	dept := new(model.SysDept)
	parent := &model.SysDept{DeptId: req.ParentId}
	ok, err := dao.DeptDao.FindOne(parent)
	if err == nil && ok {
		if parent.Status != "0" {
			return 0, errors.New("部门停用，不允许新增")
		} else {
			dept.Ancestors = parent.Ancestors + "," + gconv.String(parent.DeptId)
		}

	} else {
		return 0, errors.New("父部门不能为空")
	}

	dept.DeptName = req.DeptName
	dept.Status = req.Status
	dept.ParentId = req.ParentId
	dept.DelFlag = "0"
	dept.Email = req.Email
	dept.Leader = req.Leader
	dept.Phone = req.Phone
	dept.OrderNum = req.OrderNum
	dept.TenantId = req.TenantId
	user := UserService.GetProfile(ctx)

	if user != nil && user.UserId > 0 {
		dept.CreateBy = user.LoginName
	}

	dept.CreateTime = time.Now()

	_, err = dao.DeptDao.Insert(dept)
	return dept.DeptId, err
}

//修改保存信息
func (s *deptService) EditSave(req *model.DeptEditReq, ctx *gin.Context) (int64, error) {
	dept := &model.SysDept{DeptId: req.DeptId}
	ok, err := dao.DeptDao.FindOne(dept)
	if err != nil || !ok {
		return 0, errors.New("数据不存在")
	}
	pdept := &model.SysDept{DeptId: req.ParentId}

	ok, err = dao.DeptDao.FindOne(dept)
	if pdept != nil {
		if pdept.Status != "0" {
			return 0, errors.New("部门停用，不允许新增")
		} else {
			newAncestors := pdept.Ancestors + "," + gconv.String(pdept.DeptId)
			dao.DeptDao.UpdateDeptChildren(dept.DeptId, newAncestors, dept.Ancestors)

			dept.DeptName = req.DeptName
			dept.Status = req.Status
			dept.ParentId = req.ParentId
			dept.DelFlag = "0"
			dept.Email = req.Email
			dept.Leader = req.Leader
			dept.Phone = req.Phone
			dept.OrderNum = req.OrderNum

			user := UserService.GetProfile(ctx)

			if user != nil && user.UserId > 0 {
				dept.UpdateBy = user.LoginName
			}

			dept.UpdateTime = time.Now()

			dao.DeptDao.Update(dept)
			return 1, nil
		}

	} else {
		return 0, errors.New("父部门不能为空")
	}
}

//根据分页查询部门管理数据
func (s *deptService) SelectListAll(param *model.DeptSelectPageReq) ([]model.SysDept, error) {
	if param == nil {
		return s.SelectDeptList(0, "", "", param.TenantId)
	} else {
		return s.SelectDeptList(param.ParentId, param.DeptName, param.Status, param.TenantId)
	}
}

//删除部门管理信息
func (s *deptService) DeleteDeptById(deptId int64) int64 {
	return dao.DeptDao.DeleteDeptById(deptId)
}

//根据部门ID查询信息
func (s *deptService) SelectDeptById(deptId int64) *model.SysDeptExtend {
	deptEntity, err := dao.DeptDao.SelectDeptById(deptId)
	if err != nil {
		return nil
	}

	return deptEntity
}

//根据ID查询所有子部门
func (s *deptService) SelectChildrenDeptById(deptId int64) []*model.SysDept {
	return dao.DeptDao.SelectChildrenDeptById(deptId)
}

//加载部门列表树
func (s *deptService) SelectDeptTree(parentId int64, deptName, status string, tenantId int64) (*[]model.Ztree, error) {
	list, err := dao.DeptDao.SelectDeptList(parentId, deptName, status, tenantId)
	if err != nil {
		return nil, err
	}

	return s.InitZtree(&list, nil), nil

}

//查询部门管理数据
func (s *deptService) SelectDeptList(parentId int64, deptName, status string, tenantId int64) ([]model.SysDept, error) {
	return dao.DeptDao.SelectDeptList(parentId, deptName, status, tenantId)
}

//根据角色ID查询部门（数据权限）
func (s *deptService) RoleDeptTreeData(roleId int64, tenantId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	deptList, err := dao.DeptDao.SelectDeptList(0, "", "", tenantId)
	if err != nil {
		return nil, err
	}

	if roleId > 0 {
		roleDeptList, err := dao.DeptDao.SelectRoleDeptTree(roleId)
		if err != nil || roleDeptList == nil {
			result = s.InitZtree(&deptList, nil)
		} else {
			result = s.InitZtree(&deptList, &roleDeptList)
		}
	} else {
		result = s.InitZtree(&deptList, nil)
	}
	return result, nil
}

//对象转部门树
func (s *deptService) InitZtree(deptList *[]model.SysDept, roleDeptList *[]string) *[]model.Ztree {
	var result []model.Ztree
	isCheck := false
	if roleDeptList != nil && len(*roleDeptList) > 0 {
		isCheck = true
	}

	for i := range *deptList {
		if (*deptList)[i].Status == "0" {
			var ztree model.Ztree
			ztree.Id = (*deptList)[i].DeptId
			ztree.Pid = (*deptList)[i].ParentId
			ztree.Name = (*deptList)[i].DeptName
			ztree.Title = (*deptList)[i].DeptName
			if isCheck {
				tmp := gconv.String((*deptList)[i].DeptId) + (*deptList)[i].DeptName
				tmpcheck := false
				for j := range *roleDeptList {
					if strings.EqualFold((*roleDeptList)[j], tmp) {
						tmpcheck = true
						break
					}
				}
				ztree.Checked = tmpcheck
			}
			result = append(result, ztree)
		}
	}
	return &result
}

//查询部门是否存在用户
func (s *deptService) CheckDeptExistUser(deptId int64) bool {
	return dao.DeptDao.CheckDeptExistUser(deptId)
}

//查询部门人数
func SelectDeptCount(deptId, parentId int64) int64 {
	return dao.DeptDao.SelectDeptCount(deptId, parentId)
}

//校验部门名称是否唯一
func (s *deptService) CheckDeptNameUniqueAll(deptName string, parentId int64) string {
	dept, err := dao.DeptDao.CheckDeptNameUniqueAll(deptName, parentId)
	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 {
		return "1"
	} else {
		return "0"
	}
}

//校验部门名称是否唯一
func (s *deptService) CheckDeptNameUnique(deptName string, deptId, parentId int64) string {
	dept, err := dao.DeptDao.CheckDeptNameUniqueAll(deptName, parentId)

	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 && dept.DeptId != deptId {
		return "1"
	}
	return "0"
}
