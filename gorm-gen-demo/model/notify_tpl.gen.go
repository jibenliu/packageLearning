// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameNotifyTpl = "notify_tpl"

// NotifyTpl mapped from table <notify_tpl>
type NotifyTpl struct {
	ID      int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Channel string `gorm:"column:channel;type:varchar(32);not null" json:"channel"`
	Name    string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Content string `gorm:"column:content;type:text;not null" json:"content"`
}

// TableName NotifyTpl's table name
func (*NotifyTpl) TableName() string {
	return TableNameNotifyTpl
}
