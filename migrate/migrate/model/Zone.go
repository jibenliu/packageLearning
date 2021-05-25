package model

type VatZone struct {
	ID        uint     `gorm:"primarykey"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
	Pid       int      `gorm:"column:pid;default:0;NOT NULL" json:"pid"` // 父id
	Pkey      string   `gorm:"column:pkey" json:"pkey"`                  // 父级信息
	Name      string   `gorm:"column:name;NOT NULL" json:"name"`         // 区域名称
}
