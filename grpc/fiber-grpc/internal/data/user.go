package data

type User struct {
	Username string
	Password string
	Identity string
	// 简介
	Content string
	// 等级
	Level int
	// 头像
	Vargar string
}

func (u User) TableName() string {
	return "user"
}
