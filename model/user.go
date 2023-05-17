package model

type User struct {
	Name     string `json:"name" xorm:"pk"`
	Password string `json:"password"`
}

func (user *User) TableName() string {
	return "tb_user"
}
