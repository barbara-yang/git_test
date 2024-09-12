package handler

import (
	"entry_task/common/module"
	"entry_task/usersrv/config"
	"entry_task/usersrv/dao"
	"fmt"
)

// editNickname handle edit nickname call
func editNickname(user module.User) (module.User, error) {
	redis := dao.SessionRedis{}
	session, err := redis.Check(user.UserID)
	if err != nil {
		return module.User{}, fmt.Errorf("auth expire or no auth err:%s", err)
	}

	if session == user.Session {
		db := dao.UserDao{}
		user = db.UpdateNickname(user)
		redis.Refresh(user.UserID)
		return db.FindByID(user.UserID), nil
	}

	return module.User{}, fmt.Errorf("illegal edit")
}

// editUserProfilePic handle edit user profile pic call
func editUserProfilePic(user module.User) (module.User, error) {
	redis := dao.SessionRedis{}
	session, err := redis.Check(user.UserID)
	if err != nil {
		return module.User{}, fmt.Errorf("auth expire or no auth err:%s", err)
	}

	if session == user.Session {
		db := dao.UserDao{}
		user = db.UpdateUserProfilePic(user)
		redis.Refresh(user.UserID)
		return user, nil
	}

	return module.User{}, fmt.Errorf("illegal edit")
}

// EditUserProfilePic provide edit user profile pic method to rpc server
func EditUserProfilePic(user module.User) (module.User, error) {
	// 接口测试用
	if config.TESTFLAG {
		redis := dao.SessionRedis{}
		_, _ = redis.Check(user.UserID) //用于性能测试时计算SessionCheck时间
		if user.Session == "test-session" {
			db := dao.UserDao{}
			user = db.UpdateUserProfilePic(user)
			return user, nil
		}
	}

	return editUserProfilePic(user)
}

// EditNickname provide edit user profile pic method to rpc server
func EditNickname(user module.User) (module.User, error) {
	// 接口测试用
	if config.TESTFLAG {
		redis := dao.SessionRedis{}
		_, _ = redis.Check(user.UserID) //用于性能测试时计算SessionCheck时间
		if user.Session == "test-session" {
			db := dao.UserDao{}
			user = db.UpdateNickname(user)
			user = db.FindByID(user.UserID)
			return user, nil
		}
	}

	return editNickname(user)
}
