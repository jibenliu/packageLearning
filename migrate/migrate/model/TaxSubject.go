package model

type Taxsubject struct {
	Id            string `gorm:"column:Id;NOT NULL" json:"Id"`
	Name          string `gorm:"column:Name" json:"Name"`
	NSRSBH        string `gorm:"column:NSRSBH" json:"NSRSBH"`
	Address       string `gorm:"column:Address" json:"Address"`
	PhoneNum      string `gorm:"column:PhoneNum" json:"PhoneNum"`
	Bank          string `gorm:"column:Bank" json:"Bank"`
	BankNum       string `gorm:"column:BankNum" json:"BankNum"`
	BUGUID        string `gorm:"column:BUGUID" json:"BUGUID"`
	OrgType       int    `gorm:"column:OrgType" json:"OrgType"`
	ParentId      string `gorm:"column:ParentId" json:"ParentId"`
	LevelCode     string `gorm:"column:LevelCode" json:"LevelCode"`
	Code          string `gorm:"column:Code" json:"Code"`
	SwjgCode      string `gorm:"column:SwjgCode" json:"SwjgCode"`
	SwjgName      string `gorm:"column:SwjgName" json:"SwjgName"`
	ShortName     string `gorm:"column:ShortName" json:"ShortName"` // 税务主体简称
	OrgCode       string `gorm:"column:OrgCode" json:"OrgCode"`     // 组织编码
	VersionNumber string `gorm:"column:VersionNumber;NOT NULL" json:"VersionNumber"`
}

func (m *Taxsubject) TableName() string {
	return "ygz_taxsubject"
}
