package models

import (
	"gorm.io/gorm"
)

type Room struct {
	_                          gorm.DeletedAt
	AreaChgGUID                string  `gorm:"column:AreaChgGUID"`
	AreaStatus                 string  `gorm:"column:AreaStatus"`
	AreaStatusEnum             int     `gorm:"column:AreaStatusEnum"`
	BaBldPrice                 float32 `gorm:"column:BaBldPrice"`
	BaTnPrice                  float32 `gorm:"column:BaTnPrice"`
	BaTotal                    float32 `gorm:"column:BaTotal"`
	BcxyBldPrice               float32 `gorm:"column:BcxyBldPrice"`
	BcxyTnPrice                float32 `gorm:"column:BcxyTnPrice"`
	BcxyTotal                  float32 `gorm:"column:BcxyTotal"`
	BldArea                    float32 `gorm:"column:BldArea"`
	BldCode                    string  `gorm:"column:BldCode"`
	BldGUID                    string  `gorm:"column:BldGUID"`
	BldPrice                   float32 `gorm:"column:BldPrice"`
	BottomPriceReturnHouseLock int     `gorm:"column:BottomPriceReturnHouseLock"`
	BUGUID                     string  `gorm:"column:BUGUID"`
	CalMode                    string  `gorm:"column:CalMode"`
	CalModeEnum                int     `gorm:"column:CalModeEnum"`
	ChooseRoomGUID             string  `gorm:"column:ChooseRoomGUID"`
	ChooseRoomLockEndTime      string  `gorm:"column:ChooseRoomLockEndTime"`
	ChooseRoomLockGUID         string  `gorm:"column:ChooseRoomLockGUID"`
	ChooseRoomLockTime         string  `gorm:"column:ChooseRoomLockTime"`
	ChooseRoomTime             string  `gorm:"column:ChooseRoomTime"`
	Cx                         string  `gorm:"column:Cx"`
	DjBldPrice                 float32 `gorm:"column:DjBldPrice"`
	DjTnPrice                  float32 `gorm:"column:DjTnPrice"`
	DjTotal                    float32 `gorm:"column:DjTotal"`
	DspAreaStatus              string  `gorm:"column:DspAreaStatus"`
	DspAreaStatusEnum          int     `gorm:"column:DspAreaStatusEnum"`
	DspBldArea                 float32 `gorm:"column:DspBldArea"`
	DspTnArea                  float32 `gorm:"column:DspTnArea"`
	FloorName                  string  `gorm:"column:FloorName"`
	FloorNo                    int     `gorm:"column:FloorNo"`
	HxName                     string  `gorm:"column:HxName"`
	IsAlreadyBind              int     `gorm:"column:IsAlreadyBind"`
	IsAnnexe                   int     `gorm:"column:IsAnnexe"`
	IsAreaModify               int     `gorm:"column:IsAreaModify"`
	IsHfLock                   int     `gorm:"column:IsHfLock"`
	IsNetContract              int     `gorm:"column:IsNetContract"`
	IsTfLock                   int     `gorm:"column:IsTfLock"`
	IsTradeLock                int     `gorm:"column:IsTradeLock"`
	IsVirtualRoom              int     `gorm:"column:IsVirtualRoom"`
	IsYkkpLock                 int     `gorm:"column:IsYkkpLock"`
	JFDate                     string  `gorm:"column:JFDate"`
	Jg                         string  `gorm:"column:Jg"`
	MainRoomGUID               string  `gorm:"column:MainRoomGUID"`
	No                         string  `gorm:"column:No"`
	OpeningLock                int     `gorm:"column:OpeningLock"`
	ProjCode                   string  `gorm:"column:ProjCode"`
	ProjGUID                   string  `gorm:"column:ProjGUID"`
	Room                       string  `gorm:"column:Room"`
	RoomInfo                   string  `gorm:"column:RoomInfo"`
	RoomNo                     int     `gorm:"column:RoomNo"`
	RoomStru                   string  `gorm:"column:RoomStru"`
	ScBldArea                  float32 `gorm:"column:ScBldArea"`
	ScTnArea                   float32 `gorm:"column:ScTnArea"`
	ShortRoomInfo              string  `gorm:"column:ShortRoomInfo"`
	Status                     string  `gorm:"column:Status"`
	StatusEnum                 int     `gorm:"column:StatusEnum"`
	TfDate                     string  `gorm:"column:TfDate"`
	TnArea                     float32 `gorm:"column:TnArea"`
	TnPrice                    float32 `gorm:"column:TnPrice"`
	Total                      float32 `gorm:"column:Total"`
	TradeLocker                string  `gorm:"column:TradeLocker"`
	TradeLockTime              string  `gorm:"column:TradeLockTime"`
	Unit                       string  `gorm:"column:Unit"`
	UnitNo                     int     `gorm:"column:UnitNo"`
	Xoffset                    int     `gorm:"column:Xoffset"`
	XxCode                     string  `gorm:"column:XxCode"`
	XxDate                     string  `gorm:"column:XxDate"`
	Yoffset                    int     `gorm:"column:Yoffset"`
	YsBldArea                  float32 `gorm:"column:YsBldArea"`
	YsTnArea                   float32 `gorm:"column:YsTnArea"`
	CreatedGUID                string  `gorm:"column:CreatedGUID"`
	CreatedName                string  `gorm:"column:CreatedName"`
	ModifiedGUID               string  `gorm:"column:ModifiedGUID"`
	ModifiedName               string  `gorm:"column:ModifiedName"`
	RoomGUID                   string  `gorm:"column:RoomGUID"`
	FangPanUser                string  `gorm:"column:FangPanUser"`
	FangPanTime                string  `gorm:"column:FangPanTime"`
}

func (Room) TableName() string {
	return "s_Room"
}
