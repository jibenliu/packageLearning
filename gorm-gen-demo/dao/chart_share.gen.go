// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm-gen-demo/model"
)

func newChartShare(db *gorm.DB, opts ...gen.DOOption) chartShare {
	_chartShare := chartShare{}

	_chartShare.chartShareDo.UseDB(db, opts...)
	_chartShare.chartShareDo.UseModel(&model.ChartShare{})

	tableName := _chartShare.chartShareDo.TableName()
	_chartShare.ALL = field.NewAsterisk(tableName)
	_chartShare.ID = field.NewInt64(tableName, "id")
	_chartShare.Cluster = field.NewString(tableName, "cluster")
	_chartShare.DashboardID = field.NewInt64(tableName, "dashboard_id")
	_chartShare.DatasourceID = field.NewInt64(tableName, "datasource_id")
	_chartShare.Configs = field.NewString(tableName, "configs")
	_chartShare.CreateAt = field.NewInt64(tableName, "create_at")
	_chartShare.CreateBy = field.NewString(tableName, "create_by")

	_chartShare.fillFieldMap()

	return _chartShare
}

type chartShare struct {
	chartShareDo

	ALL          field.Asterisk
	ID           field.Int64
	Cluster      field.String
	DashboardID  field.Int64
	DatasourceID field.Int64
	Configs      field.String
	CreateAt     field.Int64
	CreateBy     field.String

	fieldMap map[string]field.Expr
}

func (c chartShare) Table(newTableName string) *chartShare {
	c.chartShareDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chartShare) As(alias string) *chartShare {
	c.chartShareDo.DO = *(c.chartShareDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chartShare) updateTableName(table string) *chartShare {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Cluster = field.NewString(table, "cluster")
	c.DashboardID = field.NewInt64(table, "dashboard_id")
	c.DatasourceID = field.NewInt64(table, "datasource_id")
	c.Configs = field.NewString(table, "configs")
	c.CreateAt = field.NewInt64(table, "create_at")
	c.CreateBy = field.NewString(table, "create_by")

	c.fillFieldMap()

	return c
}

func (c *chartShare) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chartShare) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["cluster"] = c.Cluster
	c.fieldMap["dashboard_id"] = c.DashboardID
	c.fieldMap["datasource_id"] = c.DatasourceID
	c.fieldMap["configs"] = c.Configs
	c.fieldMap["create_at"] = c.CreateAt
	c.fieldMap["create_by"] = c.CreateBy
}

func (c chartShare) clone(db *gorm.DB) chartShare {
	c.chartShareDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chartShare) replaceDB(db *gorm.DB) chartShare {
	c.chartShareDo.ReplaceDB(db)
	return c
}

type chartShareDo struct{ gen.DO }

type IChartShareDo interface {
	gen.SubQuery
	Debug() IChartShareDo
	WithContext(ctx context.Context) IChartShareDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChartShareDo
	WriteDB() IChartShareDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChartShareDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChartShareDo
	Not(conds ...gen.Condition) IChartShareDo
	Or(conds ...gen.Condition) IChartShareDo
	Select(conds ...field.Expr) IChartShareDo
	Where(conds ...gen.Condition) IChartShareDo
	Order(conds ...field.Expr) IChartShareDo
	Distinct(cols ...field.Expr) IChartShareDo
	Omit(cols ...field.Expr) IChartShareDo
	Join(table schema.Tabler, on ...field.Expr) IChartShareDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChartShareDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChartShareDo
	Group(cols ...field.Expr) IChartShareDo
	Having(conds ...gen.Condition) IChartShareDo
	Limit(limit int) IChartShareDo
	Offset(offset int) IChartShareDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChartShareDo
	Unscoped() IChartShareDo
	Create(values ...*model.ChartShare) error
	CreateInBatches(values []*model.ChartShare, batchSize int) error
	Save(values ...*model.ChartShare) error
	First() (*model.ChartShare, error)
	Take() (*model.ChartShare, error)
	Last() (*model.ChartShare, error)
	Find() ([]*model.ChartShare, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChartShare, err error)
	FindInBatches(result *[]*model.ChartShare, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ChartShare) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChartShareDo
	Assign(attrs ...field.AssignExpr) IChartShareDo
	Joins(fields ...field.RelationField) IChartShareDo
	Preload(fields ...field.RelationField) IChartShareDo
	FirstOrInit() (*model.ChartShare, error)
	FirstOrCreate() (*model.ChartShare, error)
	FindByPage(offset int, limit int) (result []*model.ChartShare, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChartShareDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chartShareDo) Debug() IChartShareDo {
	return c.withDO(c.DO.Debug())
}

func (c chartShareDo) WithContext(ctx context.Context) IChartShareDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chartShareDo) ReadDB() IChartShareDo {
	return c.Clauses(dbresolver.Read)
}

func (c chartShareDo) WriteDB() IChartShareDo {
	return c.Clauses(dbresolver.Write)
}

func (c chartShareDo) Session(config *gorm.Session) IChartShareDo {
	return c.withDO(c.DO.Session(config))
}

func (c chartShareDo) Clauses(conds ...clause.Expression) IChartShareDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chartShareDo) Returning(value interface{}, columns ...string) IChartShareDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chartShareDo) Not(conds ...gen.Condition) IChartShareDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chartShareDo) Or(conds ...gen.Condition) IChartShareDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chartShareDo) Select(conds ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chartShareDo) Where(conds ...gen.Condition) IChartShareDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chartShareDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IChartShareDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chartShareDo) Order(conds ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chartShareDo) Distinct(cols ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chartShareDo) Omit(cols ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chartShareDo) Join(table schema.Tabler, on ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chartShareDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chartShareDo) RightJoin(table schema.Tabler, on ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chartShareDo) Group(cols ...field.Expr) IChartShareDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chartShareDo) Having(conds ...gen.Condition) IChartShareDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chartShareDo) Limit(limit int) IChartShareDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chartShareDo) Offset(offset int) IChartShareDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chartShareDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChartShareDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chartShareDo) Unscoped() IChartShareDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chartShareDo) Create(values ...*model.ChartShare) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chartShareDo) CreateInBatches(values []*model.ChartShare, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chartShareDo) Save(values ...*model.ChartShare) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chartShareDo) First() (*model.ChartShare, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChartShare), nil
	}
}

func (c chartShareDo) Take() (*model.ChartShare, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChartShare), nil
	}
}

func (c chartShareDo) Last() (*model.ChartShare, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChartShare), nil
	}
}

func (c chartShareDo) Find() ([]*model.ChartShare, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChartShare), err
}

func (c chartShareDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChartShare, err error) {
	buf := make([]*model.ChartShare, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chartShareDo) FindInBatches(result *[]*model.ChartShare, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chartShareDo) Attrs(attrs ...field.AssignExpr) IChartShareDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chartShareDo) Assign(attrs ...field.AssignExpr) IChartShareDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chartShareDo) Joins(fields ...field.RelationField) IChartShareDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chartShareDo) Preload(fields ...field.RelationField) IChartShareDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chartShareDo) FirstOrInit() (*model.ChartShare, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChartShare), nil
	}
}

func (c chartShareDo) FirstOrCreate() (*model.ChartShare, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChartShare), nil
	}
}

func (c chartShareDo) FindByPage(offset int, limit int) (result []*model.ChartShare, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chartShareDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chartShareDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chartShareDo) Delete(models ...*model.ChartShare) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chartShareDo) withDO(do gen.Dao) *chartShareDo {
	c.DO = *do.(*gen.DO)
	return c
}