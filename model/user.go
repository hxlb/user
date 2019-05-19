package model

type User struct {
	Id       int32  `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Name     string `xorm:"char(60)" form:"username" json:"username"`
	Phone    string `xorm:"char(50)" form:"phone" json:"phone"`
	Password string `xorm:"char(100)" form:"password" json:"-"`
}
