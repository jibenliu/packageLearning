// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaskTplHost = "task_tpl_host"

// TaskTplHost mapped from table <task_tpl_host>
type TaskTplHost struct {
	Ii   int64  `gorm:"column:ii;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"ii"`
	ID   int64  `gorm:"column:id;type:int(10) unsigned;not null;comment:task tpl id" json:"id,string"`
	Host string `gorm:"column:host;type:varchar(128);not null;comment:ip or hostname" json:"host"`
}

// TableName TaskTplHost's table name
func (*TaskTplHost) TableName() string {
	return TableNameTaskTplHost
}