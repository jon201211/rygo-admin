package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var DictTypeDao = newDictTypeDao()

func newDictTypeDao() *dictTypeDao {
	return &dictTypeDao{}
}

type dictTypeDao struct {
}

//映射数据表
func (d *dictTypeDao) TableName() string {
	return "sys_dict_type"
}

// 插入数据
func (d *dictTypeDao) Insert(r *model.SysDictType) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *dictTypeDao) Update(r *model.SysDictType) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DictId).Update(r)
}

// 删除
func (d *dictTypeDao) Delete(r *model.SysDictType) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DictId).Delete(r)
}

//批量删除
func (d *dictTypeDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("dict_id", ids).Delete(new(model.SysDictType))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *dictTypeDao) FindOne(r *model.SysDictType) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *dictTypeDao) Find(where, order string) ([]model.SysDictType, error) {
	var list []model.SysDictType
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *dictTypeDao) FindIn(column string, args ...interface{}) ([]model.SysDictType, error) {
	var list []model.SysDictType
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *dictTypeDao) FindNotIn(column string, args ...interface{}) ([]model.SysDictType, error) {
	var list []model.SysDictType
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *dictTypeDao) SelectListByPage(param *model.DictTypeSelectPageReq) ([]model.SysDictType, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.DictName != "" {
			session.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			session.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			session.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
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

	var result []model.SysDictType
	err = session.Find(&result)
	return result, p, err
}

// 导出excel
func (d *dictTypeDao) SelectListExport(param *model.DictTypeSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.DictName != "" {
			build.Where(builder.Like{"t.dict_name", param.DictName})
		}

		if param.DictType != "" {
			build.Where(builder.Like{"t.dict_type", param.DictType})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
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
func (d *dictTypeDao) SelectListAll(param *model.DictTypeSelectPageReq) ([]model.SysDictType, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.DictName != "" {
			session.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			session.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			session.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysDictType
	err := session.Find(&result)
	return result, err
}

//校验字典类型是否唯一
func (d *dictTypeDao) CheckDictTypeUniqueAll(dictType string) (*model.SysDictType, error) {
	var entity model.SysDictType
	entity.DictType = dictType
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
