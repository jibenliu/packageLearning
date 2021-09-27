package demo

import (
	"github.com/MohamedBassem/gormgen"
	"github.com/jinzhu/gorm"
)

func search() {
	db := &gorm.DB{}
	// 查询示例：
	// The gorm way:
	var users []User
	_ = db.Where("age > ?", 20).Order("age ASC").Limit(10).Find(&users).Error

	// gormgen way
	users, _ = (&UserQueryBuilder{}).
		WhereUserName(gormgen.GreaterThanPredicate, "张三").
		OrderById(true).
		Limit(10).
		QueryAll(db)
}
