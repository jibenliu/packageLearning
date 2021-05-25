package model

import "time"

type TaxEntityMapping struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TaxEntityId int       `gorm:"column:tax_entity_id;default:0;NOT NULL" json:"tax_entity_id"`               // 税务主体id
	MappingId   string    `gorm:"column:mapping_id;NOT NULL" json:"mapping_id"`                               // 关联的对应数据表的id
	MappingType string    `gorm:"column:mapping_type;NOT NULL" json:"mapping_type"`                           // 关联类型
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`   // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:0000-00-00 00:00:00;NOT NULL" json:"update_time"` // 更新时间
}
