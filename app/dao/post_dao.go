package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"

	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var PostDao = newPostDao()

func newPostDao() *postDao {
	return &postDao{}
}

type postDao struct {
}

//映射数据表
func (d *postDao) TableName() string {
	return "sys_post"
}

// 插入数据
func (d *postDao) Insert(r *model.SysPost) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *postDao) Update(r *model.SysPost) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.PostId).Update(r)
}

// 删除
func (d *postDao) Delete(r *model.SysPost) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.PostId).Delete(r)
}

//批量删除
func (d *postDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("post_id", ids).Delete(new(model.SysPost))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *postDao) FindOne(r *model.SysPost) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *postDao) Find(where, order string) ([]model.SysPost, error) {
	var list []model.SysPost
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *postDao) FindIn(column string, args ...interface{}) ([]model.SysPost, error) {
	var list []model.SysPost
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *postDao) FindNotIn(column string, args ...interface{}) ([]model.SysPost, error) {
	var list []model.SysPost
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *postDao) SelectListByPage(param *model.PostSelectPageReq) ([]model.SysPost, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("p")

	if param != nil {
		if param.PostCode != "" {
			session.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			session.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			session.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	if param.OrderByColumn != "" {
		session.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.SysPost
	err = session.Find(&result)
	return result, p, err
}

// 导出excel
func (d *postDao) SelectListExport(param *model.PostSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.PostCode != "" {
			build.Where(builder.Like{"t.post_code", param.PostCode})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.PostName != "" {
			build.Where(builder.Like{"t.post_name", param.PostName})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

//获取所有数据
func (d *postDao) SelectListAll(param *model.PostSelectPageReq) ([]model.PostEntityFlag, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("p").Select("p.*,false as flag")
	if param != nil {

		if param.PostCode != "" {
			session.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			session.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			session.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.PostEntityFlag
	err := session.Find(&result)
	return result, err
}

//根据用户ID查询岗位
func (d *postDao) SelectPostsByUserId(userId int64) ([]model.PostEntityFlag, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("u")
	session.Join("LEFT", []string{"sys_user_post", "up"}, "u.user_id = up.user_id")
	session.Join("LEFT", []string{"sys_post", "p"}, "up.post_id = p.post_id")
	session.Where("up.user_id = ?", userId)
	session.Select("p.post_id, p.post_name, p.post_code,false as flag")
	var result []model.PostEntityFlag
	err := session.Find(&result)
	return result, err
}

//校验岗位名称是否唯一
func (d *postDao) CheckPostNameUniqueAll(postName string) (*model.SysPost, error) {
	var entity model.SysPost
	entity.PostName = postName
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

//校验岗位名称是否唯一
func (d *postDao) CheckPostCodeUniqueAll(postCode string) (*model.SysPost, error) {
	var entity model.SysPost
	entity.PostCode = postCode
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
