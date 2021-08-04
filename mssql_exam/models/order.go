package models

import (
	"gorm.io/gorm"
)

type Order struct {
	_                gorm.DeletedAt
	Addition         string  `gorm:"column:Addition"`
	AgreementNo      string  `gorm:"column:AgreementNo"`
	AjBank           string  `gorm:"column:AjBank"`
	AjTotal          float32 `gorm:"column:AjTotal"`
	AjYear           int     `gorm:"column:AjYear"`
	AreaStatus       string  `gorm:"column:AreaStatus"`
	AreaStatusEnum   int     `gorm:"column:AreaStatusEnum"`
	AuditDate        string  `gorm:"column:AuditDate"`
	AuditName        string  `gorm:"column:AuditName"`
	AuditStatus      string  `gorm:"column:AuditStatus"`
	AuditStatusEnum  int     `gorm:"column:AuditStatusEnum"`
	BaBldPrice       float32 `gorm:"column:BaBldPrice"`
	BaTnPrice        float32 `gorm:"column:BaTnPrice"`
	BaTotal          float32 `gorm:"column:BaTotal"`
	BldArea          float32 `gorm:"column:BldArea"`
	BUGUID           string  `gorm:"column:BUGUID"`
	CalMode          string  `gorm:"column:CalMode"`
	CalModeEnum      int     `gorm:"column:CalModeEnum"`
	CjBldPrice       float32 `gorm:"column:CjBldPrice"`
	CjRmbTotal       float32 `gorm:"column:CjRmbTotal"`
	CjRoomTotal      float32 `gorm:"column:CjRoomTotal"`
	CjTnPrice        float32 `gorm:"column:CjTnPrice"`
	CloseDate        string  `gorm:"column:CloseDate"`
	CloseReason      string  `gorm:"column:CloseReason"`
	CloseReasonEnum  int     `gorm:"column:CloseReasonEnum"`
	DiscntRemark     string  `gorm:"column:DiscntRemark"`
	DiscntValue      float32 `gorm:"column:DiscntValue"`
	DjBldPrice       float32 `gorm:"column:DjBldPrice"`
	DjTnPrice        float32 `gorm:"column:DjTnPrice"`
	DjTotal          float32 `gorm:"column:DjTotal"`
	Earnest          float32 `gorm:"column:Earnest"`
	FsTotal          float32 `gorm:"column:FsTotal"`
	GjjBank          string  `gorm:"column:GjjBank"`
	GjjTotal         float32 `gorm:"column:GjjTotal"`
	GjjYear          int     `gorm:"column:GjjYear"`
	IdCode           string  `gorm:"column:IdCode"`
	IsLockDisount    int     `gorm:"column:IsLockDisount"`
	IsST             int     `gorm:"column:IsST"`
	IsZxkbrht        int     `gorm:"column:IsZxkbrht"`
	JbrGUID          string  `gorm:"column:JbrGUID"`
	JbrName          string  `gorm:"column:JbrName"`
	LastSaleType     string  `gorm:"column:LastSaleType"`
	LastSaleTypeEnum int     `gorm:"column:LastSaleTypeEnum"`
	OrderBarcode     string  `gorm:"column:OrderBarcode"`
	OrderType        string  `gorm:"column:OrderType"`
	OrderTypeEnum    int     `gorm:"column:OrderTypeEnum"`
	PayFormName      string  `gorm:"column:PayFormName"`
	ProjGUID         string  `gorm:"column:ProjGUID"`
	QSDate           string  `gorm:"column:QSDate"`
	Rate             float32 `gorm:"column:Rate"`
	Remark           string  `gorm:"column:Remark"`
	RoomBldPrice     float32 `gorm:"column:RoomBldPrice"`
	RoomGUID         string  `gorm:"column:RoomGUID"`
	RoomTnPrice      float32 `gorm:"column:RoomTnPrice"`
	RoomTotal        float32 `gorm:"column:RoomTotal"`
	Status           string  `gorm:"column:Status"`
	StatusEnum       int     `gorm:"column:StatusEnum"`
	TaxAmount        float32 `gorm:"column:TaxAmount"`
	TjrAllName       string  `gorm:"column:TjrAllName"`
	TnArea           float32 `gorm:"column:TnArea"`
	TradeGUID        string  `gorm:"column:TradeGUID"`
	YqyDate          string  `gorm:"column:YqyDate"`
	YwgsDate         string  `gorm:"column:YwgsDate"`
	ZxBz             string  `gorm:"column:ZxBz"`
	ZxPrice          float32 `gorm:"column:ZxPrice"`
	ZxTotal          float32 `gorm:"column:ZxTotal"`
	Zygw             string  `gorm:"column:Zygw"`
	ZygwAllGUID      string  `gorm:"column:ZygwAllGUID"`
	CreatedGUID      string  `gorm:"column:CreatedGUID"`
	CreatedName      string  `gorm:"column:CreatedName"`
	ModifiedGUID     string  `gorm:"column:ModifiedGUID"`
	ModifiedName     string  `gorm:"column:ModifiedName"`
	OrderGUID        string  `gorm:"column:OrderGUID"`
	CstAllCardID     string  `gorm:"column:CstAllCardID"`
	CstAllName       string  `gorm:"column:CstAllName"`
	IsEnableBcxy     int     `gorm:"column:IsEnableBcxy"`
	CjxyTotal        float32 `gorm:"column:CjxyTotal"`
	RoomBzBldPrice   float32 `gorm:"column:RoomBzBldPrice"`
	RoomBzTnPrice    float32 `gorm:"column:RoomBzTnPrice"`
	RoomBzTotal      float32 `gorm:"column:RoomBzTotal"`
	CreatedTime      string  `gorm:"column:CreatedTime"`
	ModifiedTime     string  `gorm:"column:ModifiedTime"`
}

func (Order) TableName() string {
	return "s_Order"
}
