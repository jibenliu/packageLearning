package model

type TaxSubjectBank struct {
	Id           string `gorm:"column:Id;NOT NULL" json:"Id"`            // 主键
	TaxSubjectId string `gorm:"column:TaxSubjectId" json:"TaxSubjectId"` // 税务主体主键
	BankName     string `gorm:"column:BankName" json:"BankName"`         // 开户银行名称
	BankNum      string `gorm:"column:BankNum" json:"BankNum"`           // 开户银行账号
	IsDefault    int    `gorm:"column:IsDefault" json:"IsDefault"`       // 是否默认
}

func (m *TaxSubjectBank) TableName() string {
	return "ygz_taxsubjectbank"
}
