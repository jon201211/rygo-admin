package dao

import (
	"errors"
	db "rygo/app/db"
	"rygo/app/model"

	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var ConfigDao = newConfigDao()

func newConfigDao() *configDao {
	return &configDao{}
}

type configDao struct {
}

//映射数据表
func (d *configDao) TableName() string {
	return "sys_config"
}

// 插入数据
func (d *configDao) Insert(e *model.SysConfig) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(e)
}

// 更新数据
func (d *configDao) Update(e *model.SysConfig) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.ConfigId).Update(e)
}

// 删除
func (d *configDao) Delete(e *model.SysConfig) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.ConfigId).Delete(e)
}

//批量删除
func (d *configDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("config_id", ids).Delete(new(model.SysConfig))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *configDao) FindOne(e *model.SysConfig) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(e)
}

// 根据条件查询
func (d *configDao) Find(where, order string) ([]model.SysConfig, error) {
	var list []model.SysConfig
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *configDao) FindIn(column string, args ...interface{}) ([]model.SysConfig, error) {
	var list []model.SysConfig
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *configDao) FindNotIn(column string, args ...interface{}) ([]model.SysConfig, error) {
	var list []model.SysConfig
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *configDao) SelectListByPage(param *model.ConfigSelectPageReq) ([]model.SysConfig, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.ConfigName != "" {
			session.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			session.Where("t.config_type = ?", param.ConfigType)
		}

		if param.ConfigKey != "" {
			session.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	m := session.Clone()

	total, err := m.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.SysConfig
	err = session.Find(&result)
	return result, p, err
}

// 导出excel
func (d *configDao) SelectListExport(param *model.ConfigSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.ConfigName != "" {
			build.Where(builder.Like{"t.config_name", param.ConfigName})
		}

		if param.ConfigType != "" {
			build.Where(builder.Eq{"t.status": param.ConfigType})
		}

		if param.ConfigKey != "" {
			build.Where(builder.Like{"t.config_key", param.ConfigKey})
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
func (d *configDao) SelectListAll(param *model.ConfigSelectPageReq) ([]model.SysConfig, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.ConfigName != "" {
			session.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			session.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			session.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysConfig
	err := session.Find(&result)
	return result, err
}

//校验参数键名是否唯一
func (d *configDao) CheckPostCodeUniqueAll(configKey string) (*model.SysConfig, error) {
	var entity model.SysConfig
	entity.ConfigKey = configKey
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
