package handler

import (
	"entry_task/common/module"
	"entry_task/usersrv/dao"
	"fmt"
)

// GetUser provide get user method to rpc server
func GetUser(user module.User) (module.User, error) {
	redis := dao.SessionRedis{}
	db := dao.UserDao{}
	session, err := redis.Check(user.UserID)
	if err != nil {
		return module.User{}, fmt.Errorf("get user info fail auth fail err:%s", err)
	}
	if session == user.Session {
		return db.FindByID(user.UserID), nil
	}
	return module.User{}, fmt.Errorf("get user info fail auth fail ")
}
