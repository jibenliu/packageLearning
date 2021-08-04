package models

import (
	"gorm.io/gorm"
)

type Voucher struct {
	_                    gorm.DeletedAt
	AuditDate            string  `gorm:"column:AuditDate"`
	AuditName            string  `gorm:"column:AuditName"`
	BankHdAmount         float32 `gorm:"column:BankHdAmount"`
	BeforeRmbYe          float32 `gorm:"column:BeforeRmbYe"`
	BUGUID               string  `gorm:"column:BUGUID"`
	ChDate               string  `gorm:"column:ChDate"`
	ChInvoCode           string  `gorm:"column:ChInvoCode"`
	ChInvoNO             string  `gorm:"column:ChInvoNO"`
	ChJlr                string  `gorm:"column:ChJlr"`
	ClearingNo           string  `gorm:"column:ClearingNo"`
	ElectronicReceiptUrl string  `gorm:"column:ElectronicReceiptUrl"`
	ExportDate           string  `gorm:"column:ExportDate"`
	FailReason           string  `gorm:"column:FailReason"`
	GetForm              string  `gorm:"column:GetForm"`
	GetFormText          string  `gorm:"column:GetFormText"`
	Glsp                 string  `gorm:"column:Glsp"`
	HbAmount             float32 `gorm:"column:HbAmount"`
	HbDate               string  `gorm:"column:HbDate"`
	Hbr                  string  `gorm:"column:Hbr"`
	InvoCode             string  `gorm:"column:InvoCode"`
	InvoNO               string  `gorm:"column:InvoNO"`
	InvoType             string  `gorm:"column:InvoType"`
	InvoTypeEnum         int     `gorm:"column:InvoTypeEnum"`
	IsCreatedPz          int     `gorm:"column:IsCreatedPz"`
	IsExport             int     `gorm:"column:IsExport"`
	IsGenerateFail       int     `gorm:"column:IsGenerateFail"`
	IsNeedCreatePz       int     `gorm:"column:IsNeedCreatePz"`
	ItemTypeTag          int     `gorm:"column:ItemTypeTag"`
	JbDate               string  `gorm:"column:JbDate"`
	Jbr                  string  `gorm:"column:Jbr"`
	Jkr                  string  `gorm:"column:Jkr"`
	KpDate               string  `gorm:"column:KpDate"`
	Kpr                  string  `gorm:"column:Kpr"`
	kpr                  string  `gorm:"column:kpr"`
	LoanBank             string  `gorm:"column:LoanBank"`
	PosAmount            float32 `gorm:"column:PosAmount"`
	PosCode              string  `gorm:"column:PosCode"`
	PosReceiptUrl        string  `gorm:"column:PosReceiptUrl"`
	ProjGUID             string  `gorm:"column:ProjGUID"`
	Pzno                 string  `gorm:"column:Pzno"`
	ReceiptType          string  `gorm:"column:ReceiptType"`
	ReceiptTypeEnum      int     `gorm:"column:ReceiptTypeEnum"`
	Remark               string  `gorm:"column:Remark"`
	RmbAmount            float32 `gorm:"column:RmbAmount"`
	RoomGUID             string  `gorm:"column:RoomGUID"`
	RzBank               string  `gorm:"column:RzBank"`
	RzBankCode           string  `gorm:"column:RzBankCode"`
	RzDate               string  `gorm:"column:RzDate"`
	SaleGUID             string  `gorm:"column:SaleGUID"`
	SaleType             string  `gorm:"column:SaleType"`
	SaleTypeEnum         int     `gorm:"column:SaleTypeEnum"`
	SettlementForm       string  `gorm:"column:SettlementForm"`
	SkDate               string  `gorm:"column:SkDate"`
	Skdw                 string  `gorm:"column:Skdw"`
	TaxAddTel            string  `gorm:"column:TaxAddTel"`
	TaxAmount            float32 `gorm:"column:TaxAmount"`
	TaxBankAcct          string  `gorm:"column:TaxBankAcct"`
	TaxNo                string  `gorm:"column:TaxNo"`
	TaxRate              float32 `gorm:"column:TaxRate"`
	ThirdInvoiceId       float32 `gorm:"column:ThirdInvoiceId"`
	ThirdVoucherId       float32 `gorm:"column:ThirdVoucherId"`
	VoucherBarcode       float32 `gorm:"column:VoucherBarcode"`
	VouchStatus          float32 `gorm:"column:VouchStatus"`
	VouchStatusEnum      int     `gorm:"column:VouchStatusEnum"`
	VouchType            string  `gorm:"column:VouchType"`
	VouchTypeEnum        int     `gorm:"column:VouchTypeEnum"`
	ZfDate               string  `gorm:"column:ZfDate"`
	Zfr                  string  `gorm:"column:Zfr"`
	//CreatedGUID          []byte  `gorm:"column:CreatedGUID"`
	CreatedGUID          string  `gorm:"column:CreatedGUID"`
	CreatedName          string  `gorm:"column:CreatedName"`
	ModifiedGUID         string  `gorm:"column:ModifiedGUID"`
	ModifiedName         string  `gorm:"column:ModifiedName"`
	VouchGUID            string  `gorm:"column:VouchGUID"`
	CardID               string  `gorm:"column:CardID"`
	Tel                  string  `gorm:"column:Tel"`
	CreatedTime          string  `gorm:"column:CreatedTime"`
	ModifiedTime         string  `gorm:"column:ModifiedTime"`
}

func (Voucher) TableName() string {
	return "s_Voucher"
}
