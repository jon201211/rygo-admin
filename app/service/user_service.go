package service

import (
	"errors"
	"fmt"
	"rygo/app/cache"
	"rygo/app/common/session"
	"rygo/app/dao"
	"rygo/app/db"
	"rygo/app/global"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/gconv"
	"rygo/app/utils/gmd5"
	"rygo/app/utils/page"
	"rygo/app/utils/random"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

//根据主键查询用户信息
func (s *userService) SelectRecordById(id int64) (*model.SysUser, error) {
	entity := &model.SysUser{UserId: id}
	_, err := dao.UserDao.FindOne(entity)
	return entity, err
}

// 根据条件分页查询用户列表
func (s *userService) SelectRecordList(param *model.UserSelectPageReq) ([]model.UserListEntity, *page.Paging, error) {
	return dao.UserDao.SelectPageList(param)
}

// 导出excel
func (s *userService) Export(param *model.UserSelectPageReq) (string, error) {
	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	col := []string{"u.login_name", "u.user_name", "u.email", "u.phonenumber", "u.sex", "d.dept_name", "d.leader", "u.status", "u.del_flag", "u.create_by", "u.create_time", "u.remark"}
	return dao.UserDao.SelectExportList(param, head, col)
}

//新增用户
func (s *userService) AddSave(req *model.UserAddReq, ctx *gin.Context) (int64, error) {
	var user model.SysUser
	user.LoginName = req.LoginName
	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	user.CreateTime = time.Now()

	createUser := s.GetProfile(ctx)

	if createUser != nil {
		user.CreateBy = createUser.LoginName
	}

	user.DelFlag = "0"

	session := db.Instance().Engine().NewSession()
	err := session.Begin()

	_, err = session.Table(dao.UserDao.TableName()).Insert(&user)

	if err != nil || user.UserId <= 0 {
		session.Rollback()
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]model.UserPostEntity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost model.UserPostEntity
				userPost.UserId = user.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			_, err := session.Table(dao.UserPostDao.TableName()).Insert(userPosts)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]model.UserRoleEntity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole model.UserRoleEntity
				userRole.UserId = user.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			_, err := session.Table(dao.UserRoleDao.TableName()).Insert(userRoles)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		}
	}

	return user.UserId, session.Commit()
}

//新增用户
func (s *userService) EditSave(req *model.UserEditReq, ctx *gin.Context) (int64, error) {
	user := &model.SysUser{UserId: req.UserId}
	ok, err := dao.UserDao.FindOne(user)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("数据不存在")
	}

	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	user.UpdateTime = time.Now()

	updateUser := s.GetProfile(ctx)

	if updateUser != nil {
		user.UpdateBy = updateUser.LoginName
	}

	session := db.Instance().Engine().NewSession()
	tanErr := session.Begin()

	_, tanErr = session.Table(dao.UserDao.TableName()).ID(user.UserId).Update(user)

	if tanErr != nil {
		session.Rollback()
		return 0, tanErr
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]model.UserPostEntity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost model.UserPostEntity
				userPost.UserId = user.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			session.Exec("delete from sys_user_post where user_id=?", user.UserId)
			_, tanErr = session.Table(dao.UserPostDao.TableName()).Insert(userPosts)
			if tanErr != nil {
				session.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]model.UserRoleEntity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole model.UserRoleEntity
				userRole.UserId = user.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			session.Exec("delete from sys_user_role where user_id=?", user.UserId)
			_, err := session.Table(dao.UserRoleDao.TableName()).Insert(userRoles)
			if tanErr != nil {
				session.Rollback()
				return 0, err
			}
		}
	}

	return 1, session.Commit()
}

//根据主键删除用户信息
func (s *userService) DeleteRecordById(id int64) bool {
	entity := &model.SysUser{UserId: id}
	result, _ := dao.UserDao.Delete(entity)
	if result > 0 {
		return true
	}
	return false
}

//批量删除用户记录
func (s *userService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, _ := dao.UserDao.DeleteBatch(idarr...)
	dao.UserRoleDao.DeleteBatch(idarr...)
	dao.UserPostDao.DeleteBatch(idarr...)
	return result
}

//判断是否是系统管理员
func (s *userService) IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx *gin.Context) bool {
	userId, exist := ctx.Get(global.USER_ID)
	fmt.Println("IsSignedIn----------->userId:", userId)
	if exist {
		return true
	}
	return false
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) SignIn(loginnName, password string) (*model.SysUser, error) {
	//查询用户信息
	user := model.SysUser{LoginName: loginnName}
	ok, err := dao.UserDao.FindOne(&user)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("用户名或者密码错误")
	}

	//校验密码
	pwdNew := user.LoginName + password + user.Salt

	pwdNew = gmd5.MustEncryptString(pwdNew)

	if strings.Compare(user.Password, pwdNew) == -1 {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

//清空用户菜单缓存
func (s *userService) ClearMenuCache(user *model.SysUser) {
	if s.IsAdmin(user.UserId) {
		cache.Instance().Delete(global.MENU_CACHE)
	} else {
		cache.Instance().Delete(global.MENU_CACHE + gconv.String(user.UserId))
	}
}

// 用户注销
func (s *userService) SignOut(ctx *gin.Context) error {
	user := s.GetProfile(ctx)
	if user != nil {
		s.ClearMenuCache(user)
	}
	//userId := ctx.MustGet("userId").(int64)
	return nil
}

//强退用户
func (s *userService) ForceLogout(sessionId string) error {
	var tmp interface{}
	if v, ok := global.SessionList.Load(sessionId); ok {
		tmp = v
	}

	if tmp != nil {
		ctx := tmp.(*gin.Context)
		if ctx != nil {
			s.SignOut(ctx)
			global.SessionList.Delete(sessionId)
			entity := model.UserOnline{Sessionid: sessionId}
			dao.UserOnlineDao.Delete(&entity)
		}
	}

	return nil
}

// 检查账号是否符合规范,存在返回false,否则true
func (s *userService) CheckPassport(loginName string) bool {
	entity := model.SysUser{LoginName: loginName}
	if ok, err := dao.UserDao.FindOne(&entity); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (s *userService) CheckNickName(userName string) bool {
	entity := model.SysUser{UserName: userName}
	if ok, err := dao.UserDao.FindOne(&entity); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func (s *userService) CheckLoginName(loginName string) bool {
	entity := model.SysUser{LoginName: loginName}
	if ok, err := dao.UserDao.FindOne(&entity); err != nil || !ok {
		return false
	} else {
		return true
	}
}

// 获得用户信息详情
func (s *userService) GetProfile(ctx *gin.Context) *model.SysUser {

	var u = session.GetProfile(ctx)
	return u
}

//更新用户信息详情
func (s *userService) UpdateProfile(profile *model.UserProfileReq, ctx *gin.Context) error {
	user := s.GetProfile(ctx)

	if profile.UserName != "" {
		user.UserName = profile.UserName
	}

	if profile.Email != "" {
		user.Email = profile.Email
	}

	if profile.Phonenumber != "" {
		user.Phonenumber = profile.Phonenumber
	}

	if profile.Sex != "" {
		user.Sex = profile.Sex
	}

	_, err := dao.UserDao.Update(user)
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, ctx)
	return nil
}

//更新用户头像
func (s *userService) UpdateAvatar(avatar string, ctx *gin.Context) error {
	user := s.GetProfile(ctx)

	if avatar != "" {
		user.Avatar = avatar
	}

	_, err := dao.UserDao.Update(user)
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, ctx)
	return nil
}

//修改用户密码
func (s *userService) UpdatePassword(profile *model.UserPasswordReq, ctx *gin.Context) error {
	user := s.GetProfile(ctx)

	if profile.OldPassword == "" {
		return errors.New("旧密码不能为空")
	}

	if profile.NewPassword == "" {
		return errors.New("新密码不能为空")
	}

	if profile.Confirm == "" {
		return errors.New("确认密码不能为空")
	}

	if profile.NewPassword == profile.OldPassword {
		return errors.New("新旧密码不能相同")
	}

	if profile.Confirm != profile.NewPassword {
		return errors.New("确认密码不一致")
	}

	//校验密码
	token := user.LoginName + profile.OldPassword + user.Salt
	token = gmd5.MustEncryptString(token)

	if token != user.Password {
		return errors.New("原密码不正确")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	_, err := dao.UserDao.Update(user)
	if err != nil {
		return errors.New("保存数据失败")
	}

	//SaveUserToSession(user, ctx)
	return nil
}

//重置用户密码
func (s *userService) ResetPassword(params *model.UserResetPwdReq) (bool, error) {
	user := model.SysUser{UserId: params.UserId}
	if ok, err := dao.UserDao.FindOne(&user); err != nil || !ok {
		return false, errors.New("用户不存在")
	}
	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	if _, err := dao.UserDao.Update(&user); err != nil {
		return false, errors.New("保存数据失败")
	}
	return true, nil
}

//校验密码是否正确
func (s *userService) CheckPassword(user *model.SysUser, password string) bool {
	if user == nil || user.UserId <= 0 {
		return false
	}

	//校验密码
	token := user.LoginName + password + user.Salt
	token = gmd5.MustEncryptString(token)

	if strings.Compare(token, user.Password) == 0 {
		return true
	} else {
		return false
	}
}

//检查邮箱是否已使用
func (s *userService) CheckEmailUnique(userId int64, email string) bool {
	return dao.UserDao.CheckEmailUnique(userId, email)
}

//检查邮箱是否存在,存在返回true,否则false
func (s *userService) CheckEmailUniqueAll(email string) bool {
	return dao.UserDao.CheckEmailUniqueAll(email)
}

//检查手机号是否已使用,存在返回true,否则false
func (s *userService) CheckPhoneUnique(userId int64, phone string) bool {
	return dao.UserDao.CheckPhoneUnique(userId, phone)
}

//检查手机号是否已使用 ,存在返回true,否则false
func (s *userService) CheckPhoneUniqueAll(phone string) bool {
	return dao.UserDao.CheckPhoneUniqueAll(phone)
}

//根据登陆名查询用户信息
func (s *userService) SelectUserByLoginName(loginName string) (*model.SysUser, error) {
	return dao.UserDao.SelectUserByLoginName(loginName)
}

//根据手机号查询用户信息
func (s *userService) SelectUserByPhoneNumber(phonenumber string) (*model.SysUser, error) {
	return dao.UserDao.SelectUserByPhoneNumber(phonenumber)
}

// 查询已分配用户角色列表
func (s *userService) SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]model.SysUser, error) {
	return dao.UserDao.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func (s *userService) SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]model.SysUser, error) {
	return dao.UserDao.SelectUnallocatedList(roleId, loginName, phonenumber)
}
