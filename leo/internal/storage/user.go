package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(60);" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);" json:"password"`
	Identity string `gorm:"column:identity;type:varchar(100);" json:"identity"`
	
}
