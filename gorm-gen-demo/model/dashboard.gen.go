// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDashboard = "dashboard"

// Dashboard mapped from table <dashboard>
type Dashboard struct {
	ID       int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	GroupID  int64  `gorm:"column:group_id;type:bigint(20);not null;comment:busi group id" json:"group_id"`
	Name     string `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Tags     string `gorm:"column:tags;type:varchar(255);not null;comment:split by space" json:"tags"`
	Configs  string `gorm:"column:configs;type:varchar(8192);comment:dashboard variables" json:"configs"`
	CreateAt int64  `gorm:"column:create_at;type:bigint(20);not null" json:"create_at"`
	CreateBy string `gorm:"column:create_by;type:varchar(64);not null" json:"create_by"`
	UpdateAt int64  `gorm:"column:update_at;type:bigint(20);not null" json:"update_at"`
	UpdateBy string `gorm:"column:update_by;type:varchar(64);not null" json:"update_by"`
}

// TableName Dashboard's table name
func (*Dashboard) TableName() string {
	return TableNameDashboard
}