package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"entry_task/common/config"
	"entry_task/common/redis"

	"fmt"
	"strconv"
	"time"
)

// SessionRedis redis object to handle session request
type SessionRedis struct {
}

// Set creat and set session for login and return session
func (*SessionRedis) Set(userID, userName string) string {
	//value := fmt.Sprintf("%x", md5.Sum([]byte()))
	hash256 := sha256.Sum256([]byte(fmt.Sprintf("%s%s%s%s", userID, userName, config.SALT, time.Now().String())))
	value := hex.EncodeToString(hash256[:])
	result := redis.GetRedisCli().Set(userID, value, 3600*time.Second)
	if result.Err() != nil {
		fmt.Println("redis set failed:", result.Err())
		fmt.Println("redis set usersrv:", userID)
		return ""
	}
	return value
}

// Refresh refresh session when act with session and userID
func (*SessionRedis) Refresh(userID int) {
	result := redis.GetRedisCli().Expire(strconv.Itoa(userID), 3600*time.Second)
	if result.Err() != nil {
		fmt.Println("redis expire failed:", result.Err())
	}
}

// Check check if session exist
func (*SessionRedis) Check(userID int) (string, error) {
	session, err := redis.GetRedisCli().Get(strconv.Itoa(userID)).Result()
	if err != nil {
		return "", err
	}
	return session, nil
}

// Del del session when logout
func (*SessionRedis) Del(userID int) {
	redis.GetRedisCli().Del(strconv.Itoa(userID))
}
