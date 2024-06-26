// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBusiGroup = "busi_group"

// BusiGroup mapped from table <busi_group>
type BusiGroup struct {
	ID          int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Name        string `gorm:"column:name;type:varchar(191);not null" json:"name"`
	LabelEnable int64  `gorm:"column:label_enable;type:tinyint(1);not null" json:"label_enable"`
	LabelValue  string `gorm:"column:label_value;type:varchar(191);not null;comment:if label_enable: label_value can not be blank" json:"label_value"`
	CreateAt    int64  `gorm:"column:create_at;type:bigint(20);not null" json:"create_at"`
	CreateBy    string `gorm:"column:create_by;type:varchar(64);not null" json:"create_by"`
	UpdateAt    int64  `gorm:"column:update_at;type:bigint(20);not null" json:"update_at"`
	UpdateBy    string `gorm:"column:update_by;type:varchar(64);not null" json:"update_by"`
}

// TableName BusiGroup's table name
func (*BusiGroup) TableName() string {
	return TableNameBusiGroup
}
