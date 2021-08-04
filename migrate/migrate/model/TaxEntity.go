package model

type VatTaxEntity struct {
	Id        int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`         // 自增id
	Name      string `gorm:"column:name;NOT NULL" json:"name"`                       // 全称
	ZoneId    int    `gorm:"column:zone_id;default:0;NOT NULL" json:"zone_id"`       // 区域id
	Address   string `gorm:"column:address;NOT NULL" json:"address"`                 // 地址
	Mobile    string `gorm:"column:mobile;NOT NULL" json:"mobile"`                   // 电话
	TaxCode   string `gorm:"column:tax_code;NOT NULL" json:"tax_code"`               // 纳税人识别号
	HasAuth   int    `gorm:"column:has_auth;default:0;NOT NULL" json:"has_auth"`     // 是否授权
	CreatedAt int    `gorm:"column:created_at;default:0;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;default:0;NOT NULL" json:"updated_at"` // 最后更新时间
	//Banks     []VatTaxEntityBank
	//Addresses []VatTaxEntityAddressInfo
}
