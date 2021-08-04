package model

type VatTaxEntityBank struct {
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`         // 自增id
	EntityId    int    `gorm:"column:entity_id;default:0" json:"entity_id"`            // 税务主体id
	BankName    string `gorm:"column:bank_name;NOT NULL" json:"bank_name"`             // 开户行名称
	BankAccount string `gorm:"column:bank_account;NOT NULL" json:"bank_account"`       // 开户行账号
	IsDefault   int    `gorm:"column:is_default;default:0;NOT NULL" json:"is_default"` // 是否默认（0非默认1默认）
	CreatedAt   int    `gorm:"column:created_at;default:0;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt   int    `gorm:"column:updated_at;default:0;NOT NULL" json:"updated_at"` // 最后更新时间
}
