package sys

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string ``
	NickName  string ``
	AvatarUrl string ``
	Password  string ``
	Gender    int32  ``
	Telephone string ``
}
