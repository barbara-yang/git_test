package handler

import (
	"crypto/rand"
	"encoding/base64"
	"entry_task/common/module"
	"entry_task/usersrv/dao"
	"fmt"
	"log"
	"strconv"
)

// register handle use register request
func Register(username, password string) (module.User, error) {
	db := dao.UserDao{}
	user := db.FindByName(username)
	if user.Username != "" {
		return module.User{}, fmt.Errorf("user already exists")
	}

	// Hash the password
	b := make([]byte, 32)
	_, err := rand.Read(b)
	var salt string
	if err == nil {
		salt = base64.URLEncoding.EncodeToString(b)
	}

	// Create new user
	user = module.User{
		Username: username,
		Password: password,
		Salt:     salt,
	}

	// Save the new user in the database
	userID, err := db.Create(user)
	if err != nil {
		return module.User{}, fmt.Errorf("failed to create user: %s", err)
	}
	user.UserID = userID

	// Generate and set a session for the new user
	redis := dao.SessionRedis{}
	user.Session = redis.Set(strconv.Itoa(user.UserID), user.Username)

	return user, nil
}

// Login handle user login request
func Login(username, password string) (module.User, error) {
	db := dao.UserDao{}
	user := db.FindByName(username)
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password+user.Salt), []byte(password))

	if user.Password == password {
		//session存在则刷新不存在则生成新session
		var err error
		redis := dao.SessionRedis{}
		user.Session, err = redis.Check(user.UserID)
		if err != nil {
			user.Session = redis.Set(strconv.Itoa(user.UserID), user.Username)
			return user, nil
		}
		redis.Refresh(user.UserID)
		return user, nil
	}
	log.Printf("username: %s, password: %s, user: %+v\n", username, password, user)
	return module.User{}, fmt.Errorf("incorrect password")
}

// Logout handle user logout request
func Logout(user module.User) error {
	redis := dao.SessionRedis{}
	session, err := redis.Check(user.UserID)
	if err != nil {
		return fmt.Errorf("logout fail auth fail err:%s", err)
	}
	if session == user.Session {
		redis.Del(user.UserID)
		return nil
	}
	return fmt.Errorf("logout fail auth fail ")
}

// Auth handle user auth request
func Auth(user module.User) (bool, error) {
	redis := dao.SessionRedis{}
	session, err := redis.Check(user.UserID)
	if err != nil {
		return false, fmt.Errorf("auth fail err:%s", err)
	}
	if session == user.Session {
		redis.Refresh(user.UserID)
		return true, nil
	}
	return false, nil
}
