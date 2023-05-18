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

func newChart(db *gorm.DB, opts ...gen.DOOption) chart {
	_chart := chart{}

	_chart.chartDo.UseDB(db, opts...)
	_chart.chartDo.UseModel(&model.Chart{})

	tableName := _chart.chartDo.TableName()
	_chart.ALL = field.NewAsterisk(tableName)
	_chart.ID = field.NewInt64(tableName, "id")
	_chart.GroupID = field.NewInt64(tableName, "group_id")
	_chart.Configs = field.NewString(tableName, "configs")
	_chart.Weight = field.NewInt64(tableName, "weight")

	_chart.fillFieldMap()

	return _chart
}

type chart struct {
	chartDo

	ALL     field.Asterisk
	ID      field.Int64
	GroupID field.Int64 // chart group id
	Configs field.String
	Weight  field.Int64

	fieldMap map[string]field.Expr
}

func (c chart) Table(newTableName string) *chart {
	c.chartDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chart) As(alias string) *chart {
	c.chartDo.DO = *(c.chartDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chart) updateTableName(table string) *chart {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.GroupID = field.NewInt64(table, "group_id")
	c.Configs = field.NewString(table, "configs")
	c.Weight = field.NewInt64(table, "weight")

	c.fillFieldMap()

	return c
}

func (c *chart) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chart) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["id"] = c.ID
	c.fieldMap["group_id"] = c.GroupID
	c.fieldMap["configs"] = c.Configs
	c.fieldMap["weight"] = c.Weight
}

func (c chart) clone(db *gorm.DB) chart {
	c.chartDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chart) replaceDB(db *gorm.DB) chart {
	c.chartDo.ReplaceDB(db)
	return c
}

type chartDo struct{ gen.DO }

type IChartDo interface {
	gen.SubQuery
	Debug() IChartDo
	WithContext(ctx context.Context) IChartDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChartDo
	WriteDB() IChartDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChartDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChartDo
	Not(conds ...gen.Condition) IChartDo
	Or(conds ...gen.Condition) IChartDo
	Select(conds ...field.Expr) IChartDo
	Where(conds ...gen.Condition) IChartDo
	Order(conds ...field.Expr) IChartDo
	Distinct(cols ...field.Expr) IChartDo
	Omit(cols ...field.Expr) IChartDo
	Join(table schema.Tabler, on ...field.Expr) IChartDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChartDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChartDo
	Group(cols ...field.Expr) IChartDo
	Having(conds ...gen.Condition) IChartDo
	Limit(limit int) IChartDo
	Offset(offset int) IChartDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChartDo
	Unscoped() IChartDo
	Create(values ...*model.Chart) error
	CreateInBatches(values []*model.Chart, batchSize int) error
	Save(values ...*model.Chart) error
	First() (*model.Chart, error)
	Take() (*model.Chart, error)
	Last() (*model.Chart, error)
	Find() ([]*model.Chart, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Chart, err error)
	FindInBatches(result *[]*model.Chart, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Chart) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChartDo
	Assign(attrs ...field.AssignExpr) IChartDo
	Joins(fields ...field.RelationField) IChartDo
	Preload(fields ...field.RelationField) IChartDo
	FirstOrInit() (*model.Chart, error)
	FirstOrCreate() (*model.Chart, error)
	FindByPage(offset int, limit int) (result []*model.Chart, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChartDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chartDo) Debug() IChartDo {
	return c.withDO(c.DO.Debug())
}

func (c chartDo) WithContext(ctx context.Context) IChartDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chartDo) ReadDB() IChartDo {
	return c.Clauses(dbresolver.Read)
}

func (c chartDo) WriteDB() IChartDo {
	return c.Clauses(dbresolver.Write)
}

func (c chartDo) Session(config *gorm.Session) IChartDo {
	return c.withDO(c.DO.Session(config))
}

func (c chartDo) Clauses(conds ...clause.Expression) IChartDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chartDo) Returning(value interface{}, columns ...string) IChartDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chartDo) Not(conds ...gen.Condition) IChartDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chartDo) Or(conds ...gen.Condition) IChartDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chartDo) Select(conds ...field.Expr) IChartDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chartDo) Where(conds ...gen.Condition) IChartDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chartDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IChartDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chartDo) Order(conds ...field.Expr) IChartDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chartDo) Distinct(cols ...field.Expr) IChartDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chartDo) Omit(cols ...field.Expr) IChartDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chartDo) Join(table schema.Tabler, on ...field.Expr) IChartDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chartDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChartDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chartDo) RightJoin(table schema.Tabler, on ...field.Expr) IChartDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chartDo) Group(cols ...field.Expr) IChartDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chartDo) Having(conds ...gen.Condition) IChartDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chartDo) Limit(limit int) IChartDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chartDo) Offset(offset int) IChartDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chartDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChartDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chartDo) Unscoped() IChartDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chartDo) Create(values ...*model.Chart) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chartDo) CreateInBatches(values []*model.Chart, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chartDo) Save(values ...*model.Chart) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chartDo) First() (*model.Chart, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chart), nil
	}
}

func (c chartDo) Take() (*model.Chart, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chart), nil
	}
}

func (c chartDo) Last() (*model.Chart, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chart), nil
	}
}

func (c chartDo) Find() ([]*model.Chart, error) {
	result, err := c.DO.Find()
	return result.([]*model.Chart), err
}

func (c chartDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Chart, err error) {
	buf := make([]*model.Chart, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chartDo) FindInBatches(result *[]*model.Chart, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chartDo) Attrs(attrs ...field.AssignExpr) IChartDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chartDo) Assign(attrs ...field.AssignExpr) IChartDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chartDo) Joins(fields ...field.RelationField) IChartDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chartDo) Preload(fields ...field.RelationField) IChartDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chartDo) FirstOrInit() (*model.Chart, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chart), nil
	}
}

func (c chartDo) FirstOrCreate() (*model.Chart, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Chart), nil
	}
}

func (c chartDo) FindByPage(offset int, limit int) (result []*model.Chart, count int64, err error) {
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

func (c chartDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chartDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chartDo) Delete(models ...*model.Chart) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chartDo) withDO(do gen.Dao) *chartDo {
	c.DO = *do.(*gen.DO)
	return c
}
