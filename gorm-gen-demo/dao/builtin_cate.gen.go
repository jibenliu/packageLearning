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

func newBuiltinCate(db *gorm.DB, opts ...gen.DOOption) builtinCate {
	_builtinCate := builtinCate{}

	_builtinCate.builtinCateDo.UseDB(db, opts...)
	_builtinCate.builtinCateDo.UseModel(&model.BuiltinCate{})

	tableName := _builtinCate.builtinCateDo.TableName()
	_builtinCate.ALL = field.NewAsterisk(tableName)
	_builtinCate.ID = field.NewInt64(tableName, "id")
	_builtinCate.Name = field.NewString(tableName, "name")
	_builtinCate.UserID = field.NewInt64(tableName, "user_id")

	_builtinCate.fillFieldMap()

	return _builtinCate
}

type builtinCate struct {
	builtinCateDo

	ALL    field.Asterisk
	ID     field.Int64
	Name   field.String
	UserID field.Int64

	fieldMap map[string]field.Expr
}

func (b builtinCate) Table(newTableName string) *builtinCate {
	b.builtinCateDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b builtinCate) As(alias string) *builtinCate {
	b.builtinCateDo.DO = *(b.builtinCateDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *builtinCate) updateTableName(table string) *builtinCate {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Name = field.NewString(table, "name")
	b.UserID = field.NewInt64(table, "user_id")

	b.fillFieldMap()

	return b
}

func (b *builtinCate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *builtinCate) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 3)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
	b.fieldMap["user_id"] = b.UserID
}

func (b builtinCate) clone(db *gorm.DB) builtinCate {
	b.builtinCateDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b builtinCate) replaceDB(db *gorm.DB) builtinCate {
	b.builtinCateDo.ReplaceDB(db)
	return b
}

type builtinCateDo struct{ gen.DO }

type IBuiltinCateDo interface {
	gen.SubQuery
	Debug() IBuiltinCateDo
	WithContext(ctx context.Context) IBuiltinCateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBuiltinCateDo
	WriteDB() IBuiltinCateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBuiltinCateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBuiltinCateDo
	Not(conds ...gen.Condition) IBuiltinCateDo
	Or(conds ...gen.Condition) IBuiltinCateDo
	Select(conds ...field.Expr) IBuiltinCateDo
	Where(conds ...gen.Condition) IBuiltinCateDo
	Order(conds ...field.Expr) IBuiltinCateDo
	Distinct(cols ...field.Expr) IBuiltinCateDo
	Omit(cols ...field.Expr) IBuiltinCateDo
	Join(table schema.Tabler, on ...field.Expr) IBuiltinCateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBuiltinCateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBuiltinCateDo
	Group(cols ...field.Expr) IBuiltinCateDo
	Having(conds ...gen.Condition) IBuiltinCateDo
	Limit(limit int) IBuiltinCateDo
	Offset(offset int) IBuiltinCateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBuiltinCateDo
	Unscoped() IBuiltinCateDo
	Create(values ...*model.BuiltinCate) error
	CreateInBatches(values []*model.BuiltinCate, batchSize int) error
	Save(values ...*model.BuiltinCate) error
	First() (*model.BuiltinCate, error)
	Take() (*model.BuiltinCate, error)
	Last() (*model.BuiltinCate, error)
	Find() ([]*model.BuiltinCate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BuiltinCate, err error)
	FindInBatches(result *[]*model.BuiltinCate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BuiltinCate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBuiltinCateDo
	Assign(attrs ...field.AssignExpr) IBuiltinCateDo
	Joins(fields ...field.RelationField) IBuiltinCateDo
	Preload(fields ...field.RelationField) IBuiltinCateDo
	FirstOrInit() (*model.BuiltinCate, error)
	FirstOrCreate() (*model.BuiltinCate, error)
	FindByPage(offset int, limit int) (result []*model.BuiltinCate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBuiltinCateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b builtinCateDo) Debug() IBuiltinCateDo {
	return b.withDO(b.DO.Debug())
}

func (b builtinCateDo) WithContext(ctx context.Context) IBuiltinCateDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b builtinCateDo) ReadDB() IBuiltinCateDo {
	return b.Clauses(dbresolver.Read)
}

func (b builtinCateDo) WriteDB() IBuiltinCateDo {
	return b.Clauses(dbresolver.Write)
}

func (b builtinCateDo) Session(config *gorm.Session) IBuiltinCateDo {
	return b.withDO(b.DO.Session(config))
}

func (b builtinCateDo) Clauses(conds ...clause.Expression) IBuiltinCateDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b builtinCateDo) Returning(value interface{}, columns ...string) IBuiltinCateDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b builtinCateDo) Not(conds ...gen.Condition) IBuiltinCateDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b builtinCateDo) Or(conds ...gen.Condition) IBuiltinCateDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b builtinCateDo) Select(conds ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b builtinCateDo) Where(conds ...gen.Condition) IBuiltinCateDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b builtinCateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBuiltinCateDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b builtinCateDo) Order(conds ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b builtinCateDo) Distinct(cols ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b builtinCateDo) Omit(cols ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b builtinCateDo) Join(table schema.Tabler, on ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b builtinCateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b builtinCateDo) RightJoin(table schema.Tabler, on ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b builtinCateDo) Group(cols ...field.Expr) IBuiltinCateDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b builtinCateDo) Having(conds ...gen.Condition) IBuiltinCateDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b builtinCateDo) Limit(limit int) IBuiltinCateDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b builtinCateDo) Offset(offset int) IBuiltinCateDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b builtinCateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBuiltinCateDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b builtinCateDo) Unscoped() IBuiltinCateDo {
	return b.withDO(b.DO.Unscoped())
}

func (b builtinCateDo) Create(values ...*model.BuiltinCate) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b builtinCateDo) CreateInBatches(values []*model.BuiltinCate, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b builtinCateDo) Save(values ...*model.BuiltinCate) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b builtinCateDo) First() (*model.BuiltinCate, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BuiltinCate), nil
	}
}

func (b builtinCateDo) Take() (*model.BuiltinCate, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BuiltinCate), nil
	}
}

func (b builtinCateDo) Last() (*model.BuiltinCate, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BuiltinCate), nil
	}
}

func (b builtinCateDo) Find() ([]*model.BuiltinCate, error) {
	result, err := b.DO.Find()
	return result.([]*model.BuiltinCate), err
}

func (b builtinCateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BuiltinCate, err error) {
	buf := make([]*model.BuiltinCate, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b builtinCateDo) FindInBatches(result *[]*model.BuiltinCate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b builtinCateDo) Attrs(attrs ...field.AssignExpr) IBuiltinCateDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b builtinCateDo) Assign(attrs ...field.AssignExpr) IBuiltinCateDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b builtinCateDo) Joins(fields ...field.RelationField) IBuiltinCateDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b builtinCateDo) Preload(fields ...field.RelationField) IBuiltinCateDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b builtinCateDo) FirstOrInit() (*model.BuiltinCate, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BuiltinCate), nil
	}
}

func (b builtinCateDo) FirstOrCreate() (*model.BuiltinCate, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BuiltinCate), nil
	}
}

func (b builtinCateDo) FindByPage(offset int, limit int) (result []*model.BuiltinCate, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b builtinCateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b builtinCateDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b builtinCateDo) Delete(models ...*model.BuiltinCate) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *builtinCateDo) withDO(do gen.Dao) *builtinCateDo {
	b.DO = *do.(*gen.DO)
	return b
}