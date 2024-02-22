package data

type User struct {
	Username string `gorm:"column:username;type:varchar(100);" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);" json:"password"`
	Identity string `gorm:"column:identity;type:varchar(100);" json:"identity"`
	// 简介
	Content string `gorm:"column:content;type:varchar(300);" json:"content"`
	// 等级
	Level int `gorm:"column:level;type:int(3);" json:"level"`
	// 头像
	Vargar string `gorm:"column:vargar;type:varchar(200);" json:"vargar"`
	// email
	Email string `gorm:"column:email;type:varchar(100);" json:"email"`
}

func (u User) TableName() string {
	return "user"
}
