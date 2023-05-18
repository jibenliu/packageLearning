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

func newNotifyTpl(db *gorm.DB, opts ...gen.DOOption) notifyTpl {
	_notifyTpl := notifyTpl{}

	_notifyTpl.notifyTplDo.UseDB(db, opts...)
	_notifyTpl.notifyTplDo.UseModel(&model.NotifyTpl{})

	tableName := _notifyTpl.notifyTplDo.TableName()
	_notifyTpl.ALL = field.NewAsterisk(tableName)
	_notifyTpl.ID = field.NewInt64(tableName, "id")
	_notifyTpl.Channel = field.NewString(tableName, "channel")
	_notifyTpl.Name = field.NewString(tableName, "name")
	_notifyTpl.Content = field.NewString(tableName, "content")

	_notifyTpl.fillFieldMap()

	return _notifyTpl
}

type notifyTpl struct {
	notifyTplDo

	ALL     field.Asterisk
	ID      field.Int64
	Channel field.String
	Name    field.String
	Content field.String

	fieldMap map[string]field.Expr
}

func (n notifyTpl) Table(newTableName string) *notifyTpl {
	n.notifyTplDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notifyTpl) As(alias string) *notifyTpl {
	n.notifyTplDo.DO = *(n.notifyTplDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notifyTpl) updateTableName(table string) *notifyTpl {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.Channel = field.NewString(table, "channel")
	n.Name = field.NewString(table, "name")
	n.Content = field.NewString(table, "content")

	n.fillFieldMap()

	return n
}

func (n *notifyTpl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notifyTpl) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 4)
	n.fieldMap["id"] = n.ID
	n.fieldMap["channel"] = n.Channel
	n.fieldMap["name"] = n.Name
	n.fieldMap["content"] = n.Content
}

func (n notifyTpl) clone(db *gorm.DB) notifyTpl {
	n.notifyTplDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n notifyTpl) replaceDB(db *gorm.DB) notifyTpl {
	n.notifyTplDo.ReplaceDB(db)
	return n
}

type notifyTplDo struct{ gen.DO }

type INotifyTplDo interface {
	gen.SubQuery
	Debug() INotifyTplDo
	WithContext(ctx context.Context) INotifyTplDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INotifyTplDo
	WriteDB() INotifyTplDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INotifyTplDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INotifyTplDo
	Not(conds ...gen.Condition) INotifyTplDo
	Or(conds ...gen.Condition) INotifyTplDo
	Select(conds ...field.Expr) INotifyTplDo
	Where(conds ...gen.Condition) INotifyTplDo
	Order(conds ...field.Expr) INotifyTplDo
	Distinct(cols ...field.Expr) INotifyTplDo
	Omit(cols ...field.Expr) INotifyTplDo
	Join(table schema.Tabler, on ...field.Expr) INotifyTplDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INotifyTplDo
	RightJoin(table schema.Tabler, on ...field.Expr) INotifyTplDo
	Group(cols ...field.Expr) INotifyTplDo
	Having(conds ...gen.Condition) INotifyTplDo
	Limit(limit int) INotifyTplDo
	Offset(offset int) INotifyTplDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INotifyTplDo
	Unscoped() INotifyTplDo
	Create(values ...*model.NotifyTpl) error
	CreateInBatches(values []*model.NotifyTpl, batchSize int) error
	Save(values ...*model.NotifyTpl) error
	First() (*model.NotifyTpl, error)
	Take() (*model.NotifyTpl, error)
	Last() (*model.NotifyTpl, error)
	Find() ([]*model.NotifyTpl, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NotifyTpl, err error)
	FindInBatches(result *[]*model.NotifyTpl, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NotifyTpl) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INotifyTplDo
	Assign(attrs ...field.AssignExpr) INotifyTplDo
	Joins(fields ...field.RelationField) INotifyTplDo
	Preload(fields ...field.RelationField) INotifyTplDo
	FirstOrInit() (*model.NotifyTpl, error)
	FirstOrCreate() (*model.NotifyTpl, error)
	FindByPage(offset int, limit int) (result []*model.NotifyTpl, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INotifyTplDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n notifyTplDo) Debug() INotifyTplDo {
	return n.withDO(n.DO.Debug())
}

func (n notifyTplDo) WithContext(ctx context.Context) INotifyTplDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n notifyTplDo) ReadDB() INotifyTplDo {
	return n.Clauses(dbresolver.Read)
}

func (n notifyTplDo) WriteDB() INotifyTplDo {
	return n.Clauses(dbresolver.Write)
}

func (n notifyTplDo) Session(config *gorm.Session) INotifyTplDo {
	return n.withDO(n.DO.Session(config))
}

func (n notifyTplDo) Clauses(conds ...clause.Expression) INotifyTplDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n notifyTplDo) Returning(value interface{}, columns ...string) INotifyTplDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n notifyTplDo) Not(conds ...gen.Condition) INotifyTplDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n notifyTplDo) Or(conds ...gen.Condition) INotifyTplDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n notifyTplDo) Select(conds ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n notifyTplDo) Where(conds ...gen.Condition) INotifyTplDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n notifyTplDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INotifyTplDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n notifyTplDo) Order(conds ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n notifyTplDo) Distinct(cols ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n notifyTplDo) Omit(cols ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n notifyTplDo) Join(table schema.Tabler, on ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n notifyTplDo) LeftJoin(table schema.Tabler, on ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n notifyTplDo) RightJoin(table schema.Tabler, on ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n notifyTplDo) Group(cols ...field.Expr) INotifyTplDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n notifyTplDo) Having(conds ...gen.Condition) INotifyTplDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n notifyTplDo) Limit(limit int) INotifyTplDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n notifyTplDo) Offset(offset int) INotifyTplDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n notifyTplDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INotifyTplDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n notifyTplDo) Unscoped() INotifyTplDo {
	return n.withDO(n.DO.Unscoped())
}

func (n notifyTplDo) Create(values ...*model.NotifyTpl) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n notifyTplDo) CreateInBatches(values []*model.NotifyTpl, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n notifyTplDo) Save(values ...*model.NotifyTpl) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n notifyTplDo) First() (*model.NotifyTpl, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTpl), nil
	}
}

func (n notifyTplDo) Take() (*model.NotifyTpl, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTpl), nil
	}
}

func (n notifyTplDo) Last() (*model.NotifyTpl, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTpl), nil
	}
}

func (n notifyTplDo) Find() ([]*model.NotifyTpl, error) {
	result, err := n.DO.Find()
	return result.([]*model.NotifyTpl), err
}

func (n notifyTplDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NotifyTpl, err error) {
	buf := make([]*model.NotifyTpl, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n notifyTplDo) FindInBatches(result *[]*model.NotifyTpl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n notifyTplDo) Attrs(attrs ...field.AssignExpr) INotifyTplDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n notifyTplDo) Assign(attrs ...field.AssignExpr) INotifyTplDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n notifyTplDo) Joins(fields ...field.RelationField) INotifyTplDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n notifyTplDo) Preload(fields ...field.RelationField) INotifyTplDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n notifyTplDo) FirstOrInit() (*model.NotifyTpl, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTpl), nil
	}
}

func (n notifyTplDo) FirstOrCreate() (*model.NotifyTpl, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTpl), nil
	}
}

func (n notifyTplDo) FindByPage(offset int, limit int) (result []*model.NotifyTpl, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n notifyTplDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n notifyTplDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n notifyTplDo) Delete(models ...*model.NotifyTpl) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *notifyTplDo) withDO(do gen.Dao) *notifyTplDo {
	n.DO = *do.(*gen.DO)
	return n
}
