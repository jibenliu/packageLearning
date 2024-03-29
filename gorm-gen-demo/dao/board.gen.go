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

func newBoard(db *gorm.DB, opts ...gen.DOOption) board {
	_board := board{}

	_board.boardDo.UseDB(db, opts...)
	_board.boardDo.UseModel(&model.Board{})

	tableName := _board.boardDo.TableName()
	_board.ALL = field.NewAsterisk(tableName)
	_board.ID = field.NewInt64(tableName, "id")
	_board.GroupID = field.NewInt64(tableName, "group_id")
	_board.Name = field.NewString(tableName, "name")
	_board.Ident = field.NewString(tableName, "ident")
	_board.Tags = field.NewString(tableName, "tags")
	_board.Public = field.NewInt64(tableName, "public")
	_board.BuiltIn = field.NewInt64(tableName, "built_in")
	_board.Hide = field.NewInt64(tableName, "hide")
	_board.CreateAt = field.NewInt64(tableName, "create_at")
	_board.CreateBy = field.NewString(tableName, "create_by")
	_board.UpdateAt = field.NewInt64(tableName, "update_at")
	_board.UpdateBy = field.NewString(tableName, "update_by")

	_board.fillFieldMap()

	return _board
}

type board struct {
	boardDo

	ALL      field.Asterisk
	ID       field.Int64
	GroupID  field.Int64 // busi group id
	Name     field.String
	Ident    field.String
	Tags     field.String // split by space
	Public   field.Int64  // 0:false 1:true
	BuiltIn  field.Int64  // 0:false 1:true
	Hide     field.Int64  // 0:false 1:true
	CreateAt field.Int64
	CreateBy field.String
	UpdateAt field.Int64
	UpdateBy field.String

	fieldMap map[string]field.Expr
}

func (b board) Table(newTableName string) *board {
	b.boardDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b board) As(alias string) *board {
	b.boardDo.DO = *(b.boardDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *board) updateTableName(table string) *board {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.GroupID = field.NewInt64(table, "group_id")
	b.Name = field.NewString(table, "name")
	b.Ident = field.NewString(table, "ident")
	b.Tags = field.NewString(table, "tags")
	b.Public = field.NewInt64(table, "public")
	b.BuiltIn = field.NewInt64(table, "built_in")
	b.Hide = field.NewInt64(table, "hide")
	b.CreateAt = field.NewInt64(table, "create_at")
	b.CreateBy = field.NewString(table, "create_by")
	b.UpdateAt = field.NewInt64(table, "update_at")
	b.UpdateBy = field.NewString(table, "update_by")

	b.fillFieldMap()

	return b
}

func (b *board) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *board) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 12)
	b.fieldMap["id"] = b.ID
	b.fieldMap["group_id"] = b.GroupID
	b.fieldMap["name"] = b.Name
	b.fieldMap["ident"] = b.Ident
	b.fieldMap["tags"] = b.Tags
	b.fieldMap["public"] = b.Public
	b.fieldMap["built_in"] = b.BuiltIn
	b.fieldMap["hide"] = b.Hide
	b.fieldMap["create_at"] = b.CreateAt
	b.fieldMap["create_by"] = b.CreateBy
	b.fieldMap["update_at"] = b.UpdateAt
	b.fieldMap["update_by"] = b.UpdateBy
}

func (b board) clone(db *gorm.DB) board {
	b.boardDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b board) replaceDB(db *gorm.DB) board {
	b.boardDo.ReplaceDB(db)
	return b
}

type boardDo struct{ gen.DO }

type IBoardDo interface {
	gen.SubQuery
	Debug() IBoardDo
	WithContext(ctx context.Context) IBoardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBoardDo
	WriteDB() IBoardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBoardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBoardDo
	Not(conds ...gen.Condition) IBoardDo
	Or(conds ...gen.Condition) IBoardDo
	Select(conds ...field.Expr) IBoardDo
	Where(conds ...gen.Condition) IBoardDo
	Order(conds ...field.Expr) IBoardDo
	Distinct(cols ...field.Expr) IBoardDo
	Omit(cols ...field.Expr) IBoardDo
	Join(table schema.Tabler, on ...field.Expr) IBoardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBoardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBoardDo
	Group(cols ...field.Expr) IBoardDo
	Having(conds ...gen.Condition) IBoardDo
	Limit(limit int) IBoardDo
	Offset(offset int) IBoardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBoardDo
	Unscoped() IBoardDo
	Create(values ...*model.Board) error
	CreateInBatches(values []*model.Board, batchSize int) error
	Save(values ...*model.Board) error
	First() (*model.Board, error)
	Take() (*model.Board, error)
	Last() (*model.Board, error)
	Find() ([]*model.Board, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Board, err error)
	FindInBatches(result *[]*model.Board, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Board) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBoardDo
	Assign(attrs ...field.AssignExpr) IBoardDo
	Joins(fields ...field.RelationField) IBoardDo
	Preload(fields ...field.RelationField) IBoardDo
	FirstOrInit() (*model.Board, error)
	FirstOrCreate() (*model.Board, error)
	FindByPage(offset int, limit int) (result []*model.Board, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBoardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b boardDo) Debug() IBoardDo {
	return b.withDO(b.DO.Debug())
}

func (b boardDo) WithContext(ctx context.Context) IBoardDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b boardDo) ReadDB() IBoardDo {
	return b.Clauses(dbresolver.Read)
}

func (b boardDo) WriteDB() IBoardDo {
	return b.Clauses(dbresolver.Write)
}

func (b boardDo) Session(config *gorm.Session) IBoardDo {
	return b.withDO(b.DO.Session(config))
}

func (b boardDo) Clauses(conds ...clause.Expression) IBoardDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b boardDo) Returning(value interface{}, columns ...string) IBoardDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b boardDo) Not(conds ...gen.Condition) IBoardDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b boardDo) Or(conds ...gen.Condition) IBoardDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b boardDo) Select(conds ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b boardDo) Where(conds ...gen.Condition) IBoardDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b boardDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBoardDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b boardDo) Order(conds ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b boardDo) Distinct(cols ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b boardDo) Omit(cols ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b boardDo) Join(table schema.Tabler, on ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b boardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBoardDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b boardDo) RightJoin(table schema.Tabler, on ...field.Expr) IBoardDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b boardDo) Group(cols ...field.Expr) IBoardDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b boardDo) Having(conds ...gen.Condition) IBoardDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b boardDo) Limit(limit int) IBoardDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b boardDo) Offset(offset int) IBoardDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b boardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBoardDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b boardDo) Unscoped() IBoardDo {
	return b.withDO(b.DO.Unscoped())
}

func (b boardDo) Create(values ...*model.Board) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b boardDo) CreateInBatches(values []*model.Board, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b boardDo) Save(values ...*model.Board) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b boardDo) First() (*model.Board, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Board), nil
	}
}

func (b boardDo) Take() (*model.Board, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Board), nil
	}
}

func (b boardDo) Last() (*model.Board, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Board), nil
	}
}

func (b boardDo) Find() ([]*model.Board, error) {
	result, err := b.DO.Find()
	return result.([]*model.Board), err
}

func (b boardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Board, err error) {
	buf := make([]*model.Board, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b boardDo) FindInBatches(result *[]*model.Board, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b boardDo) Attrs(attrs ...field.AssignExpr) IBoardDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b boardDo) Assign(attrs ...field.AssignExpr) IBoardDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b boardDo) Joins(fields ...field.RelationField) IBoardDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b boardDo) Preload(fields ...field.RelationField) IBoardDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b boardDo) FirstOrInit() (*model.Board, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Board), nil
	}
}

func (b boardDo) FirstOrCreate() (*model.Board, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Board), nil
	}
}

func (b boardDo) FindByPage(offset int, limit int) (result []*model.Board, count int64, err error) {
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

func (b boardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b boardDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b boardDo) Delete(models ...*model.Board) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *boardDo) withDO(do gen.Dao) *boardDo {
	b.DO = *do.(*gen.DO)
	return b
}
