package db

import (
	"crypto/rand"
	"encoding/base64"
	"entry_task/common/config"
	"entry_task/common/module"
	"entry_task/usersrv/dao"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// CreateUser Method for deployer
func CreateUser(c chan bool) {
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go createUser(i)
	}
	defer fmt.Printf("users created")
	wg.Wait()
	c <- true

}

// createUser method creat user
func createUser(incr int) {
	defer wg.Done()
	userDao := dao.UserDao{}
	for i := incr * 500000; i < (incr+1)*500000; i++ {
		userID := i + 1
		b := make([]byte, 32)
		_, err := rand.Read(b)
		var salt string
		if err == nil {
			salt = base64.URLEncoding.EncodeToString(b)
		}

		//hash, err := bcrypt.GenerateFromPassword([]byte("test"+strconv.Itoa(userID)+salt), bcrypt.MinCost) //加密处理
		//if err != nil {
		//	fmt.Println(err)
		//}
		password := string("test" + strconv.Itoa(userID))
		createTime := time.Now()
		//没写nickname
		user := module.User{UserID: userID, Username: fmt.Sprintf("test%d", userID), ProfilePic: config.DEFAULTPIC, Password: password, Salt: salt, CreateTime: createTime, UpdateTime: createTime}
		err1 := userDao.Insert(user)
		if err1 != nil {
			return
		}

	}

}
