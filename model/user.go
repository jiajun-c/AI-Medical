package model

type User struct {
	Username string `json:"username" xorm:"pk"`
	Password string `json:"password"`
}

func (user *User) TableName() string {
	return "tb_user"
}
