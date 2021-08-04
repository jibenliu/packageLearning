package models

import (
	"gorm.io/gorm"
)

type GetIn struct {
	_                 gorm.DeletedAt
	AfterRmbYe        float32 `gorm:"column:AfterRmbYe"`
	AfterYe           float32 `gorm:"column:AfterYe"`
	Amount            float32 `gorm:"column:Amount"`
	BeforeRmbYe       float32 `gorm:"column:BeforeRmbYe"`
	BeforeYe          float32 `gorm:"column:BeforeYe"`
	Bz                string  `gorm:"column:Bz"`
	ExRate            float32 `gorm:"column:ExRate"`
	FsettlCode        string  `gorm:"column:FsettlCode"`
	FsettleNo         string  `gorm:"column:FsettleNo"`
	GetDate           string  `gorm:"column:GetDate"`
	GetForm           string  `gorm:"column:GetForm"`
	GetFormText       string  `gorm:"column:GetFormText"`
	InSequence        int     `gorm:"column:InSequence"`
	IsEnabledCurrency int     `gorm:"column:IsEnabledCurrency"`
	IsSysCx           int     `gorm:"column:IsSysCx"`
	ItemName          string  `gorm:"column:ItemName"`
	ItemNameGUID      string  `gorm:"column:ItemNameGUID"`
	ItemType          string  `gorm:"column:ItemType"`
	ItemTypeGUID      string  `gorm:"column:ItemTypeGUID"`
	OutAmount         float32 `gorm:"column:OutAmount"`
	PosAmount         float32 `gorm:"column:PosAmount"`
	PosBankCard       string  `gorm:"column:PosBankCard"`
	PosCode           string  `gorm:"column:PosCode"`
	PosTerminal       string  `gorm:"column:PosTerminal"`
	ProjGUID          string  `gorm:"column:ProjGUID"`
	RmbAmount         float32 `gorm:"column:RmbAmount"`
	RoomGUID          string  `gorm:"column:RoomGUID"`
	RzBank            string  `gorm:"column:RzBank"`
	RzBankCode        string  `gorm:"column:RzBankCode"`
	SaleGUID          string  `gorm:"column:SaleGUID"`
	SaleType          string  `gorm:"column:SaleType"`
	SaleTypeEnum      int     `gorm:"column:SaleTypeEnum"`
	Sequence          int     `gorm:"column:Sequence"`
	Status            string  `gorm:"column:Status"`
	StatusEnum        int     `gorm:"column:StatusEnum"`
	TaxAmount         float32 `gorm:"column:TaxAmount"`
	TaxRate           float32 `gorm:"column:TaxRate"`
	ThirdGetinId      string  `gorm:"column:ThirdGetinId"`
	VouchGUID         string  `gorm:"column:VouchGUID"`
	ZzdSequence       int     `gorm:"column:ZzdSequence"`
	CreatedGUID       string  `gorm:"column:CreatedGUID"`
	CreatedName       string  `gorm:"column:CreatedName"`
	GetinGUID         string  `gorm:"column:GetinGUID"`
	ModifiedGUID      string  `gorm:"column:ModifiedGUID"`
	ModifiedName      string  `gorm:"column:ModifiedName"`
	CreatedTime       string  `gorm:"column:CreatedTime"`
	ModifiedTime      string  `gorm:"column:ModifiedTime"`
}

func (GetIn) TableName() string {
	return "s_Getin"
}
