// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSsoConfig = "sso_config"

// SsoConfig mapped from table <sso_config>
type SsoConfig struct {
	ID      int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Name    string `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Content string `gorm:"column:content;type:text;not null" json:"content"`
}

// TableName SsoConfig's table name
func (*SsoConfig) TableName() string {
	return TableNameSsoConfig
}
