package models

type Fee struct {
	Amount         float32 `gorm:"column:Amount"`
	ApplyDate      string  `gorm:"column:ApplyDate"`
	ApplyName      string  `gorm:"column:ApplyName"`
	AuditState     string  `gorm:"column:AuditState"`
	AuditStateEnum int     `gorm:"column:AuditStateEnum"`
	BUGUID         string  `gorm:"column:BUGUID"`
	Bz             string  `gorm:"column:Bz"`
	BzEnum         int     `gorm:"column:BzEnum"`
	DsAmount       float32 `gorm:"column:DsAmount"`
	ExRate         float32 `gorm:"column:ExRate"`
	Flag           string  `gorm:"column:Flag"`
	IsChg          int     `gorm:"column:IsChg"`
	IsSetFkDate    int     `gorm:"column:IsSetFkDate"`
	IsUseWorkFlow  int     `gorm:"column:IsUseWorkFlow"`
	ItemName       string  `gorm:"column:ItemName"`
	ItemNameGuid   string  `gorm:"column:ItemNameGuid"`
	ItemType       string  `gorm:"column:ItemType"`
	ItemTypeGuid   string  `gorm:"column:ItemTypeGuid"`
	JmLateFee      float32 `gorm:"column:JmLateFee"`
	JmReason       float32 `gorm:"column:JmReason"`
	lastDate       string  `gorm:"column:lastDate"`
	OutAmount      float32 `gorm:"column:OutAmount"`
	OutRmbAmount   float32 `gorm:"column:OutRmbAmount"`
	PayLagQty      int     `gorm:"column:PayLagQty"`
	PayLagUnit     string  `gorm:"column:PayLagUnit"`
	PayLagUnitEnum int     `gorm:"column:PayLagUnitEnum"`
	ProjGUID       string  `gorm:"column:ProjGUID"`
	Remark         string  `gorm:"column:Remark"`
	RmbAmount      float32 `gorm:"column:RmbAmount"`
	RmbDsAmount    float32 `gorm:"column:RmbDsAmount"`
	RmbYe          float32 `gorm:"column:RmbYe"`
	Sequence       int     `gorm:"column:Sequence"`
	TaxAmount      float32 `gorm:"column:TaxAmount"`
	TradeGUID      string  `gorm:"column:TradeGUID"`
	Ye             float32 `gorm:"column:Ye"`
	CreatedGUID    string  `gorm:"column:CreatedGUID"`
	CreatedName    string  `gorm:"column:CreatedName"`
	CreatedTime    string  `gorm:"column:CreatedTime"`
	FeeGUID        string  `gorm:"column:FeeGUID"`
	ModifiedGUID   string  `gorm:"column:ModifiedGUID"`
	ModifiedName   string  `gorm:"column:ModifiedName"`
	ModifiedTime   string  `gorm:"column:ModifiedTime"`
	AuditDate      string  `gorm:"column:AuditDate"`
	AuditName      string  `gorm:"column:AuditName"`
}

func (Fee) TableName() string {
	return "s_Fee"
}
