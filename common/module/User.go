package module

import (
	"time"
)

// User is module to save userInfo
type User struct {
	UserID     int
	Username   string
	Nickname   string
	Password   string
	ProfilePic string
	Salt string
	Session    string `sql:"-"`
	CreateTime time.Time
	UpdateTime time.Time
}
