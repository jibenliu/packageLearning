// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTarget = "target"

// Target mapped from table <target>
type Target struct {
	ID       int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	GroupID  int64  `gorm:"column:group_id;type:bigint(20);not null;comment:busi group id" json:"group_id"`
	Ident    string `gorm:"column:ident;type:varchar(191);not null;comment:target id" json:"ident"`
	Note     string `gorm:"column:note;type:varchar(255);not null;comment:append to alert event as field" json:"note"`
	Tags     string `gorm:"column:tags;type:varchar(512);not null;comment:append to series data as tags, split by space, append external space at suffix" json:"tags"`
	UpdateAt int64  `gorm:"column:update_at;type:bigint(20);not null" json:"update_at"`
}

// TableName Target's table name
func (*Target) TableName() string {
	return TableNameTarget
}
