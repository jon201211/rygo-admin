package service

import (
	"errors"
	"rygo/app/dao"
	"rygo/app/db"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/gconv"
	"rygo/app/utils/page"
	"time"

	"github.com/gin-gonic/gin"
)

var RoleService = newRoleService()

func newRoleService() *roleService {
	return &roleService{}
}

type roleService struct {
}

//根据主键查询数据
func (s *roleService) SelectRecordById(id int64) (*model.RoleEntity, error) {
	entity := &model.RoleEntity{RoleId: id}
	_, err := dao.RoleDao.FindOne(entity)
	return entity, err
}

//根据条件查询数据
func (s *roleService) SelectRecordAll(params *model.RoleSelectPageReq) ([]model.RoleEntityFlag, error) {
	return dao.RoleDao.SelectListAll(params)
}

//根据条件分页查询数据
func (s *roleService) SelectRecordPage(params *model.RoleSelectPageReq) ([]model.RoleEntity, *page.Paging, error) {
	return dao.RoleDao.SelectListPage(params)
}

//根据主键删除数据
func (s *roleService) DeleteRecordById(id int64) bool {
	entity := &model.RoleEntity{RoleId: id}
	rs, _ := dao.RoleDao.Delete(entity)
	if rs > 0 {
		return true
	}
	return false
}

//添加数据
func (s *roleService) AddSave(req *model.RoleAddReq, ctx *gin.Context) (int64, error) {

	role := new(model.RoleEntity)
	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.CreateTime = time.Now()
	role.CreateBy = ""
	role.DelFlag = "0"
	role.DataScope = "1"

	user := UserService.GetProfile(ctx)

	if user != nil {
		role.CreateBy = user.LoginName
	}

	session := db.Instance().Engine().NewSession()

	err := session.Begin()

	_, err = session.Insert(role)

	if err != nil {
		session.Rollback()
		return 0, err
	}

	if req.MenuIds != "" {
		menus := convert.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.RoleMenuEntity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp model.RoleMenuEntity
					tmp.RoleId = role.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				_, err := session.Table(dao.RoleMenuDao.TableName()).Insert(roleMenus)
				if err != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	err = session.Commit()
	return role.RoleId, err
}

//修改数据
func (s *roleService) EditSave(req *model.RoleEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.RoleEntity{RoleId: req.RoleId}
	ok, err := dao.RoleDao.FindOne(entity)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}
	entity.RoleName = req.RoleName
	entity.RoleKey = req.RoleKey
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.CreateBy = user.LoginName
	}

	session := db.Instance().Engine().NewSession()

	pErr := session.Begin()

	_, pErr = session.Table(dao.RoleDao.TableName()).ID(entity.RoleId).Update(entity)

	if pErr != nil {
		session.Rollback()
		return 0, pErr
	}

	if req.MenuIds != "" {
		menus := convert.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.RoleMenuEntity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp model.RoleMenuEntity
					tmp.RoleId = entity.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				_, pErr = session.Exec("delete from sys_role_menu where role_id=?", entity.RoleId)
				if pErr != nil {
					session.Rollback()
					return 0, pErr
				}
				_, pErr = session.Table(dao.RoleMenuDao.TableName()).Insert(roleMenus)
				if pErr != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, session.Commit()
}

//保存数据权限
func (s *roleService) AuthDataScope(req *model.RoleDataScopeReq, ctx *gin.Context) (int64, error) {
	entity := &model.RoleEntity{RoleId: req.RoleId}
	ok, err := dao.RoleDao.FindOne(entity)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	if req.DataScope != "" {
		entity.DataScope = req.DataScope
	}

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.UpdateBy = user.LoginName
	}
	entity.UpdateTime = time.Now()

	session := db.Instance().Engine().NewSession()
	tanErr := session.Begin()

	_, tanErr = session.Table(dao.RoleDao.TableName()).ID(entity.RoleId).Update(entity)

	if tanErr != nil {
		session.Rollback()
		return 0, err
	}

	if req.DeptIds != "" {
		deptids := convert.ToInt64Array(req.DeptIds, ",")
		if len(deptids) > 0 {
			roleDepts := make([]model.RoleDeptEntity, 0)
			for i := range deptids {
				if deptids[i] > 0 {
					var tmp model.RoleDeptEntity
					tmp.RoleId = entity.RoleId
					tmp.DeptId = deptids[i]
					roleDepts = append(roleDepts, tmp)
				}
			}
			if len(roleDepts) > 0 {
				session.Exec("delete from  sys_role_dept where role_id=?", entity.RoleId)
				_, tanErr := session.Table(dao.RoleDeptDao.TableName()).Insert(roleDepts)
				if tanErr != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, session.Commit()

}

//批量删除数据记录
func (s *roleService) DeleteRecordByIds(ids string) int64 {
	idArr := convert.ToInt64Array(ids, ",")
	result, _ := dao.RoleDao.DeleteBatch(idArr...)
	return result
}

// 导出excel
func (s *roleService) Export(param *model.RoleSelectPageReq) (string, error) {
	head := []string{"角色ID", "角色名称", "权限字符串", "显示顺序", "数据范围", "角色状态"}
	col := []string{"role_id", "role_name", "role_key", "role_sort", "data_scope", "status"}
	return dao.RoleDao.SelectListExport(param, head, col)
}

//根据用户ID查询角色
func (s *roleService) SelectRoleContactVo(userId int64) ([]model.RoleEntityFlag, error) {
	var paramsPost *model.RoleSelectPageReq
	roleAll, err := dao.RoleDao.SelectListAll(paramsPost)

	if err != nil || roleAll == nil {
		return nil, errors.New("未查询到岗位数据")
	}

	userRole, err := dao.RoleDao.SelectRoleContactVo(userId)

	if err != nil || userRole == nil {
		return nil, errors.New("未查询到用户岗位数据")
	} else {
		for i := range roleAll {
			for j := range userRole {
				if userRole[j].RoleId == roleAll[i].RoleId {
					roleAll[i].Flag = true
					break
				}
			}
		}
	}
	return roleAll, nil
}

//批量选择用户授权
func (s *roleService) InsertAuthUsers(roleId int64, userIds string) int64 {
	idarr := convert.ToInt64Array(userIds, ",")
	var roleUserList []model.UserRoleEntity
	for _, str := range idarr {
		var tmp model.UserRoleEntity
		tmp.UserId = str
		tmp.RoleId = roleId
		roleUserList = append(roleUserList, tmp)
	}

	rs, err := db.Instance().Engine().Table(dao.UserRoleDao.TableName()).Insert(roleUserList)
	if err != nil {
		return 0
	}
	return rs
}

//取消授权用户角色
func (s *roleService) DeleteUserRoleInfo(userId, roleId int64) int64 {
	entity := &model.UserRoleEntity{UserId: userId, RoleId: roleId}
	rs, _ := dao.UserRoleDao.Delete(entity)
	return rs
}

//批量取消授权用户角色
func (s *roleService) DeleteUserRoleInfos(roleId int64, ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")

	idStr := ""
	for _, item := range idarr {
		if item > 0 {
			if idStr != "" {
				idStr += "," + gconv.String(item)
			} else {
				idStr = gconv.String(item)
			}
		}
	}

	rs, err := db.Instance().Engine().Exec("delete from sys_user_role where role_id=? and user_id in (?)", roleId, idStr)
	if err != nil {
		return 0
	}
	nums, err := rs.RowsAffected()
	if err != nil {
		return 0
	}
	return nums
}

//检查角色名是否唯一
func (s *roleService) CheckRoleNameUniqueAll(roleName string) string {
	entity, err := dao.RoleDao.CheckRoleNameUniqueAll(roleName)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 {
		return "1"
	}
	return "0"
}

//检查角色键是否唯一
func (s *roleService) CheckRoleKeyUniqueAll(roleKey string) string {
	entity, err := dao.RoleDao.CheckRoleKeyUniqueAll(roleKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 {
		return "1"
	}
	return "0"
}

//检查角色名是否唯一
func (s *roleService) CheckRoleNameUnique(roleName string, roleId int64) string {
	entity, err := dao.RoleDao.CheckRoleNameUniqueAll(roleName)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 && entity.RoleId != roleId {
		return "1"
	}
	return "0"
}

//检查角色键是否唯一
func (s *roleService) CheckRoleKeyUnique(roleKey string, roleId int64) string {
	entity, err := dao.RoleDao.CheckRoleKeyUniqueAll(roleKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 && entity.RoleId != roleId {
		return "1"
	}
	return "0"
}

//判断是否是管理员
func (s *roleService) IsAdmin(id int64) bool {
	if id == 1 {
		return true
	} else {
		return false
	}
}

//校验角色是否允许操作
func (s *roleService) CheckRoleAllowed(id int64) bool {
	if s.IsAdmin(id) {
		return false
	} else {
		return true
	}
}
