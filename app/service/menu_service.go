package service

import (
	"errors"
	cache "rygo/app/cache"
	"rygo/app/dao"
	"rygo/app/global"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/gconv"
	"rygo/app/utils/page"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var MenuService = newMenuService()

func newMenuService() *menuService {
	return &menuService{}
}

type menuService struct {
}

//根据主键查询数据
func (s *menuService) SelectRecordById(id int64) (*model.MenuEntityExtend, error) {
	return dao.MenuDao.SelectRecordById(id)
}

//根据条件查询数据
func (s *menuService) SelectListAll(params *model.MenuSelectPageReq) ([]model.MenuEntity, error) {
	return dao.MenuDao.SelectListAll(params)
}

//根据条件分页查询数据
func (s *menuService) SelectListPage(params *model.MenuSelectPageReq) (*[]model.MenuEntity, *page.Paging, error) {
	return dao.MenuDao.SelectListPage(params)
}

//根据主键删除数据
func (s *menuService) DeleteRecordById(id int64) bool {
	entity := &model.MenuEntity{MenuId: id}
	rs, err := dao.MenuDao.Delete(entity)
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//添加数据
func (s *menuService) AddSave(req *model.MenuAddReq, ctx *gin.Context) (int64, error) {

	var entity model.MenuEntity
	entity.MenuName = req.MenuName
	entity.Visible = req.Visible
	entity.ParentId = req.ParentId
	entity.Remark = ""
	entity.MenuType = req.MenuType
	entity.Url = req.Url
	entity.Perms = req.Perms
	entity.Target = req.Target
	entity.Icon = req.Icon
	entity.OrderNum = req.OrderNum
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.MenuDao.Insert(&entity)
	return entity.MenuId, err
}

//修改数据
func (s *menuService) EditSave(req *model.MenuEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.MenuEntity{MenuId: req.MenuId}
	ok, err := dao.MenuDao.FindOne(entity)

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("角色不存在")
	}

	entity.MenuName = req.MenuName
	entity.Visible = req.Visible
	entity.ParentId = req.ParentId
	entity.Remark = ""
	entity.MenuType = req.MenuType
	entity.Url = req.Url
	entity.Perms = req.Perms
	entity.Target = req.Target
	entity.Icon = req.Icon
	entity.OrderNum = req.OrderNum
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return dao.MenuDao.Update(entity)
}

//批量删除数据记录
func (s *menuService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.MenuDao.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}
	return result
}

//加载所有菜单列表树
func (s *menuService) MenuTreeData(userId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	menuList, err := s.SelectMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}
	result, err = s.InitZtree(menuList, nil, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//获取用户的菜单数据
func (s *menuService) SelectMenuNormalByUser(userId int64) (*[]model.MenuEntityExtend, error) {
	if UserService.IsAdmin(userId) {
		return s.SelectMenuNormalAll()
	} else {
		return s.SelectMenusByUserId(gconv.String(userId))
	}
}

//获取管理员菜单数据
func (s *menuService) SelectMenuNormalAll() (*[]model.MenuEntityExtend, error) {

	//从缓存读取
	ctx := cache.Instance()
	tmp, f := ctx.Get(global.MENU_CACHE)

	if f && tmp != nil {
		rs, ok := tmp.([]model.MenuEntityExtend)
		if ok {
			return &rs, nil
		}
	}

	//从数据库中读取
	var result []model.MenuEntityExtend
	result, err := dao.MenuDao.SelectMenuNormalAll()

	if err != nil {
		return nil, err
	}

	for i := range result {
		chilrens := s.getMenuChildPerms(result, result[i].MenuId)

		for j := range chilrens {
			chilrens2 := s.getMenuChildPerms(result, chilrens[j].MenuId)
			chilrens[j].Children = chilrens2

			if chilrens[j].Target == "" {
				chilrens[j].Target = "menuItem"
			}
			if chilrens[j].Url == "" {
				chilrens[j].Url = "#"
			}
		}

		if chilrens != nil {
			result[i].Children = chilrens

			if result[i].ParentId != 0 {
				if result[i].Target == "" {
					result[i].Target = "menuItem"
				}

				if result[i].Url == "" {
					result[i].Url = "#"
				}
			}

		}
	}

	//存入缓存
	cache.Instance().Set(global.MENU_CACHE, result, time.Hour)
	return &result, nil
}

//根据用户ID读取菜单数据
func (s *menuService) SelectMenusByUserId(userId string) (*[]model.MenuEntityExtend, error) {
	//从缓存读取
	tmp, have := cache.Instance().Get(global.MENU_CACHE + userId)

	if have && tmp != nil {
		rs, ok := tmp.([]model.MenuEntityExtend)
		if ok {
			return &rs, nil
		}
	}

	//从数据库中读取
	var result []model.MenuEntityExtend
	result, err := dao.MenuDao.SelectMenusByUserId(userId)

	if err != nil {
		return nil, err
	}

	for i := range result {
		chilrens := s.getMenuChildPerms(result, result[i].MenuId)

		for j := range chilrens {
			chilrens2 := s.getMenuChildPerms(result, chilrens[j].MenuId)
			chilrens[j].Children = chilrens2
			if chilrens[j].Target == "" {
				chilrens[j].Target = "menuItem"
			}
			if chilrens[j].Url == "" {
				chilrens[j].Url = "#"
			}
		}

		if chilrens != nil {
			result[i].Children = chilrens
			if result[i].ParentId != 0 {
				if result[i].Target == "" {
					result[i].Target = "menuItem"
				}

				if result[i].Url == "" {
					result[i].Url = "#"
				}
			} else {
				if result[i].Url == "" || result[i].Url == "#" {
					result[i].Target = ""
				}
				if result[i].Url == "" {
					result[i].Url = "#"
				}
			}
		}
	}

	//存入缓存
	cache.Instance().Set(global.MENU_CACHE+userId, result, time.Hour)
	return &result, nil
}

//根据父id获取子菜单
func (s *menuService) getMenuChildPerms(menus []model.MenuEntityExtend, parentId int64) []model.MenuEntityExtend {
	if menus == nil {
		return nil
	}

	var result []model.MenuEntityExtend
	//得到一级菜单
	for i := range menus {
		if menus[i].ParentId == parentId && (menus[i].MenuType == "M" || menus[i].MenuType == "C") {
			if menus[i].Target == "" {
				menus[i].Target = "menuItem"
			}

			if menus[i].Url == "" {
				menus[i].Url = "#"
			}

			result = append(result, menus[i])
		}
	}

	return result
}

//检查菜单名是否唯一
func (s *menuService) CheckMenuNameUniqueAll(menuName string, parentId int64) string {
	entity, err := dao.MenuDao.CheckMenuNameUniqueAll(menuName, parentId)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查菜单名是否唯一
func (s *menuService) CheckMenuNameUnique(menuName string, menuId, parentId int64) string {
	entity, err := dao.MenuDao.CheckMenuNameUniqueAll(menuName, parentId)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 && entity.MenuId != menuId {
		return "1"
	}
	return "0"
}

//检查权限键是否唯一
func (s *menuService) CheckPermsUniqueAll(perms string) string {
	entity, err := dao.MenuDao.CheckPermsUniqueAll(perms)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查权限键是否唯一
func (s *menuService) CheckPermsUnique(perms string, menuId int64) string {
	entity, err := dao.MenuDao.CheckPermsUniqueAll(perms)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 && entity.MenuId != menuId {
		return "1"
	}
	return "0"
}

//根据角色ID查询菜单
func (s *menuService) RoleMenuTreeData(roleId, userId int64) (*[]model.Ztree, error) {
	var result *[]model.Ztree
	menuList, err := s.SelectMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}

	if roleId > 0 {
		roleMenuList, err := dao.MenuDao.SelectMenuTree(roleId)
		if err != nil || roleMenuList == nil {
			result, err = s.InitZtree(menuList, nil, true)
		} else {
			result, err = s.InitZtree(menuList, &roleMenuList, true)
		}
	} else {
		result, err = s.InitZtree(menuList, nil, true)
	}

	return result, nil
}

//对象转菜单树
func (s *menuService) InitZtree(menuList *[]model.MenuEntityExtend, roleMenuList *[]string, permsFlag bool) (*[]model.Ztree, error) {
	var result []model.Ztree
	isCheck := false
	if roleMenuList != nil && len(*roleMenuList) > 0 {
		isCheck = true
	}

	for _, obj := range *menuList {
		var ztree model.Ztree
		ztree.Title = obj.MenuName
		ztree.Id = obj.MenuId
		ztree.Name = s.transMenuName(obj.MenuName, permsFlag)
		ztree.Pid = obj.ParentId
		if isCheck {
			tmp := gconv.String(obj.MenuId) + obj.Perms
			tmpcheck := false
			for j := range *roleMenuList {
				if strings.Compare((*roleMenuList)[j], tmp) == 0 {
					tmpcheck = true
					break
				}
			}
			ztree.Checked = tmpcheck
		}
		result = append(result, ztree)
	}

	return &result, nil
}

func (s *menuService) transMenuName(menuName string, permsFlag bool) string {
	if permsFlag {
		return "<font color=\"#888\">&nbsp;&nbsp;&nbsp;" + menuName + "</font>"
	} else {
		return menuName
	}
}
