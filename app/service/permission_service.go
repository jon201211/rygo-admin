package service

import (
	"html/template"

	"rygo/app/model"

	"rygo/app/utils/gconv"
	"strings"
)

var PermissionService = newPermissionService()

func newPermissionService() *permissionService {
	return &permissionService{}
}

type permissionService struct {
}

//根据用户id和权限字符串判断是否输出控制按钮
func (s *permissionService) GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) template.HTML {

	result := s.HasPermi(u, permission)

	htmlstr := ""
	if result == "" {
		htmlstr = `<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
                    <i class="` + iclassName + `"></i> ` + text + `
                </a>`
	}

	return template.HTML(htmlstr)
}

//根据用户id和权限字符串判断是否有此权限
func (s *permissionService) HasPermi(u interface{}, permission string) string {
	if u == nil {
		return "disabled"
	}

	uid := gconv.Int64(u)

	if uid <= 0 {
		return "disabled"
	}
	//获取权限信息
	var menus *[]model.MenuEntityExtend
	if UserService.IsAdmin(uid) {
		menus, _ = MenuService.SelectMenuNormalAll()
	} else {
		menus, _ = MenuService.SelectMenusByUserId(gconv.String(uid))
	}

	if menus != nil && len(*menus) > 0 {
		for i := range *menus {
			if strings.EqualFold((*menus)[i].Perms, permission) {
				return ""
			}
		}
	}

	return "disabled"
}
