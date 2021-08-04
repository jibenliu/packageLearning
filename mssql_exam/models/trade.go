package models

import (
	"gorm.io/gorm"
)

type Trade struct {
	_                gorm.DeletedAt
	BcJzStatus       string `gorm:"column:BcJzStatus"`
	BcJzStatusEnum   int    `gorm:"column:BcJzStatusEnum"`
	BUGUID           string `gorm:"column:BUGUID"`
	BuyerAllNames    string `gorm:"column:BuyerAllNames"`
	CloseReason      string `gorm:"column:CloseReason"`
	CloseReasonEnum  int    `gorm:"column:CloseReasonEnum"`
	ContractQsDate   string `gorm:"column:ContractQsDate"`
	ContractYwgsDate string `gorm:"column:ContractYwgsDate"`
	JzStatus         string `gorm:"column:JzStatus"`
	JzStatusEnum     int    `gorm:"column:JzStatusEnum"`
	LastGjDate       string `gorm:"column:LastGjDate"`
	ProjGUID         string `gorm:"column:ProjGUID"`
	RGOrderGUID      string `gorm:"column:RGOrderGUID"`
	RGOrderQsDate    string `gorm:"column:RGOrderQsDate"`
	RGOrderType      string `gorm:"column:RGOrderType"`
	RGOrderTypeEnum  int    `gorm:"column:RGOrderTypeEnum"`
	RGOrderYwgsDate  string `gorm:"column:RGOrderYwgsDate"`
	RoomGUID         string `gorm:"column:RoomGUID"`
	RoomStatus       string `gorm:"column:RoomStatus"`
	RoomStatusEnum   int    `gorm:"column:RoomStatusEnum"`
	TradeStatus      string `gorm:"column:TradeStatus"`
	TradeStatusEnum  int    `gorm:"column:TradeStatusEnum"`
	ZcContractDate   string `gorm:"column:ZcContractDate"`
	ZcOrderDate      string `gorm:"column:ZcOrderDate"`
	ZcOrderGUID      string `gorm:"column:ZcOrderGUID"`
	CreatedGUID      string `gorm:"column:CreatedGUID"`
	CreatedName      string `gorm:"column:CreatedName"`
	ModifiedGUID     string `gorm:"column:ModifiedGUID"`
	ModifiedName     string `gorm:"column:ModifiedName"`
	TradeGUID        string `gorm:"column:TradeGUID"`
	BuyerAllCardIds  string `gorm:"column:BuyerAllCardIds"`
	IsExistDelayPay  int    `gorm:"column:IsExistDelayPay"`
	CreatedTime      string `gorm:"column:CreatedTime"`
	ModifiedTime     string `gorm:"column:ModifiedTime"`
}

func (Trade) TableName() string {
	return "s_Trade"
}
