// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID   int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Name string `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Note string `gorm:"column:note;type:varchar(255);not null" json:"note"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
