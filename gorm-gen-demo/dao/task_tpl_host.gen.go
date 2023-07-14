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

func newTaskTplHost(db *gorm.DB, opts ...gen.DOOption) taskTplHost {
	_taskTplHost := taskTplHost{}

	_taskTplHost.taskTplHostDo.UseDB(db, opts...)
	_taskTplHost.taskTplHostDo.UseModel(&model.TaskTplHost{})

	tableName := _taskTplHost.taskTplHostDo.TableName()
	_taskTplHost.ALL = field.NewAsterisk(tableName)
	_taskTplHost.Ii = field.NewInt64(tableName, "ii")
	_taskTplHost.ID = field.NewInt64(tableName, "id")
	_taskTplHost.Host = field.NewString(tableName, "host")

	_taskTplHost.fillFieldMap()

	return _taskTplHost
}

type taskTplHost struct {
	taskTplHostDo

	ALL  field.Asterisk
	Ii   field.Int64
	ID   field.Int64  // task tpl id
	Host field.String // ip or hostname

	fieldMap map[string]field.Expr
}

func (t taskTplHost) Table(newTableName string) *taskTplHost {
	t.taskTplHostDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskTplHost) As(alias string) *taskTplHost {
	t.taskTplHostDo.DO = *(t.taskTplHostDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskTplHost) updateTableName(table string) *taskTplHost {
	t.ALL = field.NewAsterisk(table)
	t.Ii = field.NewInt64(table, "ii")
	t.ID = field.NewInt64(table, "id")
	t.Host = field.NewString(table, "host")

	t.fillFieldMap()

	return t
}

func (t *taskTplHost) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskTplHost) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["ii"] = t.Ii
	t.fieldMap["id"] = t.ID
	t.fieldMap["host"] = t.Host
}

func (t taskTplHost) clone(db *gorm.DB) taskTplHost {
	t.taskTplHostDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t taskTplHost) replaceDB(db *gorm.DB) taskTplHost {
	t.taskTplHostDo.ReplaceDB(db)
	return t
}

type taskTplHostDo struct{ gen.DO }

type ITaskTplHostDo interface {
	gen.SubQuery
	Debug() ITaskTplHostDo
	WithContext(ctx context.Context) ITaskTplHostDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITaskTplHostDo
	WriteDB() ITaskTplHostDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITaskTplHostDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITaskTplHostDo
	Not(conds ...gen.Condition) ITaskTplHostDo
	Or(conds ...gen.Condition) ITaskTplHostDo
	Select(conds ...field.Expr) ITaskTplHostDo
	Where(conds ...gen.Condition) ITaskTplHostDo
	Order(conds ...field.Expr) ITaskTplHostDo
	Distinct(cols ...field.Expr) ITaskTplHostDo
	Omit(cols ...field.Expr) ITaskTplHostDo
	Join(table schema.Tabler, on ...field.Expr) ITaskTplHostDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITaskTplHostDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITaskTplHostDo
	Group(cols ...field.Expr) ITaskTplHostDo
	Having(conds ...gen.Condition) ITaskTplHostDo
	Limit(limit int) ITaskTplHostDo
	Offset(offset int) ITaskTplHostDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskTplHostDo
	Unscoped() ITaskTplHostDo
	Create(values ...*model.TaskTplHost) error
	CreateInBatches(values []*model.TaskTplHost, batchSize int) error
	Save(values ...*model.TaskTplHost) error
	First() (*model.TaskTplHost, error)
	Take() (*model.TaskTplHost, error)
	Last() (*model.TaskTplHost, error)
	Find() ([]*model.TaskTplHost, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskTplHost, err error)
	FindInBatches(result *[]*model.TaskTplHost, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TaskTplHost) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITaskTplHostDo
	Assign(attrs ...field.AssignExpr) ITaskTplHostDo
	Joins(fields ...field.RelationField) ITaskTplHostDo
	Preload(fields ...field.RelationField) ITaskTplHostDo
	FirstOrInit() (*model.TaskTplHost, error)
	FirstOrCreate() (*model.TaskTplHost, error)
	FindByPage(offset int, limit int) (result []*model.TaskTplHost, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITaskTplHostDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t taskTplHostDo) Debug() ITaskTplHostDo {
	return t.withDO(t.DO.Debug())
}

func (t taskTplHostDo) WithContext(ctx context.Context) ITaskTplHostDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskTplHostDo) ReadDB() ITaskTplHostDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskTplHostDo) WriteDB() ITaskTplHostDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskTplHostDo) Session(config *gorm.Session) ITaskTplHostDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskTplHostDo) Clauses(conds ...clause.Expression) ITaskTplHostDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskTplHostDo) Returning(value interface{}, columns ...string) ITaskTplHostDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskTplHostDo) Not(conds ...gen.Condition) ITaskTplHostDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskTplHostDo) Or(conds ...gen.Condition) ITaskTplHostDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskTplHostDo) Select(conds ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskTplHostDo) Where(conds ...gen.Condition) ITaskTplHostDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskTplHostDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITaskTplHostDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t taskTplHostDo) Order(conds ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskTplHostDo) Distinct(cols ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskTplHostDo) Omit(cols ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskTplHostDo) Join(table schema.Tabler, on ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskTplHostDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskTplHostDo) RightJoin(table schema.Tabler, on ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskTplHostDo) Group(cols ...field.Expr) ITaskTplHostDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskTplHostDo) Having(conds ...gen.Condition) ITaskTplHostDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskTplHostDo) Limit(limit int) ITaskTplHostDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskTplHostDo) Offset(offset int) ITaskTplHostDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskTplHostDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskTplHostDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskTplHostDo) Unscoped() ITaskTplHostDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskTplHostDo) Create(values ...*model.TaskTplHost) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskTplHostDo) CreateInBatches(values []*model.TaskTplHost, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskTplHostDo) Save(values ...*model.TaskTplHost) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskTplHostDo) First() (*model.TaskTplHost, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTplHost), nil
	}
}

func (t taskTplHostDo) Take() (*model.TaskTplHost, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTplHost), nil
	}
}

func (t taskTplHostDo) Last() (*model.TaskTplHost, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTplHost), nil
	}
}

func (t taskTplHostDo) Find() ([]*model.TaskTplHost, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskTplHost), err
}

func (t taskTplHostDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskTplHost, err error) {
	buf := make([]*model.TaskTplHost, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskTplHostDo) FindInBatches(result *[]*model.TaskTplHost, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskTplHostDo) Attrs(attrs ...field.AssignExpr) ITaskTplHostDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskTplHostDo) Assign(attrs ...field.AssignExpr) ITaskTplHostDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskTplHostDo) Joins(fields ...field.RelationField) ITaskTplHostDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskTplHostDo) Preload(fields ...field.RelationField) ITaskTplHostDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskTplHostDo) FirstOrInit() (*model.TaskTplHost, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTplHost), nil
	}
}

func (t taskTplHostDo) FirstOrCreate() (*model.TaskTplHost, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTplHost), nil
	}
}

func (t taskTplHostDo) FindByPage(offset int, limit int) (result []*model.TaskTplHost, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t taskTplHostDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskTplHostDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskTplHostDo) Delete(models ...*model.TaskTplHost) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskTplHostDo) withDO(do gen.Dao) *taskTplHostDo {
	t.DO = *do.(*gen.DO)
	return t
}