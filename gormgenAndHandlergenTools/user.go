package demo


//go:generate gormgen -structs User -output user_gen.go
type User struct {
	Id uint `gorm:"column:id" db:"id" json:"id" form:"id"` //主键
	UserName string `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"` //用户名
	NickName string `gorm:"column:nick_name" db:"nick_name" json:"nick_name" form:"nick_name"` //昵称
	Mobile string `gorm:"column:mobile" db:"mobile" json:"mobile" form:"mobile"` //手机号
	IsDeleted int8 `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted" form:"is_deleted"` //是否删除 1:是  -1:否
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"` //更新时间
}