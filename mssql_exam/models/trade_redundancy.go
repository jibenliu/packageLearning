package models

import (
	"gorm.io/gorm"
)

type TradeReDundancy struct {
	_                   gorm.DeletedAt
	BcAmount            float32 `gorm:"column:BcAmount"`
	BcRate              float32 `gorm:"column:BcRate"`
	DhpYe               float32 `gorm:"column:DhpYe"`
	HouseAmount         float32 `gorm:"column:HouseAmount"`
	PayedRate           float32 `gorm:"column:PayedRate"`
	TradeGUID           string  `gorm:"column:TradeGUID"`
	YjBcAmount          float32 `gorm:"column:YjBcAmount"`
	YsEarnest           float32 `gorm:"column:YsEarnest"`
	YsEarnestYe         float32 `gorm:"column:YsEarnestYe"`
	YsHouseAmount       float32 `gorm:"column:YsHouseAmount"`
	YsHouseFundAmount   float32 `gorm:"column:YsHouseFundAmount"`
	YsMortgageAmount    float32 `gorm:"column:YsMortgageAmount"`
	CreatedGUID         string  `gorm:"column:CreatedGUID"`
	CreatedName         string  `gorm:"column:CreatedName"`
	ModifiedGUID        string  `gorm:"column:ModifiedGUID"`
	ModifiedName        string  `gorm:"column:ModifiedName"`
	TradeRedundancyGUID string  `gorm:"column:TradeRedundancyGUID"`
	CreatedTime         string  `gorm:"column:CreatedTime"`
	ModifiedTime        string  `gorm:"column:ModifiedTime"`
}

func (TradeReDundancy) TableName() string {
	return "s_TradeReDundancy"
}
