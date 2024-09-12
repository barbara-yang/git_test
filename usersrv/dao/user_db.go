package dao

import (
	"entry_task/common/config"
	"entry_task/common/db"
	"entry_task/common/module"
	"fmt"
)

// UserDao object to handle user
type UserDao struct {
}

// Insert User
func (*UserDao) Insert(user module.User) error {
	db1 := db.GetMysqlClient()
	query := "INSERT INTO users (user_id, username, profile_pic, password, salt, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := db1.Exec(query, user.UserID, user.Username, user.ProfilePic, user.Password, user.Salt, user.CreateTime, user.UpdateTime)
	if err != nil {
		return fmt.Errorf("insert error:%v", err)
	}
	return nil
}

// 注册用户名和密码
func (*UserDao) Create(user module.User) (int, error) {
	db1 := db.GetMysqlClient()
	var err error
	_, err = db1.Exec("INSERT INTO users (username, password, salt) VALUES (?, ?, ?)", user.Username, user.Password, user.Salt)
	if err != nil {
		fmt.Printf("insert err:%v\n", err)
		return 0, err
	}
	//
	return user.UserID, nil
}

// FindByName find user by using username
func (*UserDao) FindByName(username string) module.User {
	user := new(module.User)
	var err error
	err = db.GetMysqlClient().QueryRow(
		"SELECT user_id, username,password, salt, profile_pic, create_time, update_time FROM users WHERE username = ?",
		username,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.Salt,
		&user.ProfilePic,
		&user.CreateTime,
		&user.UpdateTime,
	)
	if err != nil {
		fmt.Printf("find by name err:%v\n", err)
	}
	return *user
}

// FindByID find user by using UserId
func (*UserDao) FindByID(userID int) module.User {
	user := new(module.User)
	var err error
	err = db.GetMysqlClient().QueryRow(
		"SELECT user_id, username,nickname,password, salt, profile_pic, create_time, update_time FROM users WHERE user_id = ?",
		userID,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Nickname,
		&user.Password,
		&user.Salt,
		&user.ProfilePic,
		&user.CreateTime,
		&user.UpdateTime,
	)

	if err != nil {
		//db.GetMysqlClient().Where("user_id = ?", userID).First(&user)
		fmt.Printf("find by ID error: %v\n", err)
	}
	return *user
}

// UpdateUserProfilePic update user profilePic path when it not "" ,set profilePic as config.DEFAULTPIC when it equal ""
func (*UserDao) UpdateUserProfilePic(user module.User) module.User {
	if user.ProfilePic == "" {
		user.ProfilePic = config.DEFAULTPIC
	}

	var err error
	_, err = db.GetMysqlClient().Exec(
		"UPDATE users SET profile_pic = ? WHERE user_id = ?",
		user.ProfilePic,
		user.UserID,
	)
	if err != nil {
		fmt.Printf("update err:%v\n", err)
	}
	return user

	//err := d.db.Model(&usersrv).Where("user_id = ?",usersrv.UserId).Updates(usersrv).Error
}

// UpdateNickname user nickname
func (*UserDao) UpdateNickname(user module.User) module.User {

	var err error
	_, err = db.GetMysqlClient().Exec(
		"UPDATE users SET nickname = ? WHERE user_id = ?",
		user.Nickname,
		user.UserID,
	)

	if err != nil {
		fmt.Printf("update err:%v\n", err)
	}
	return user

}

// CreateTable to create Users table
func (*UserDao) CreateTable() {
	sql := `CREATE TABLE IF NOT EXISTS users(
	user_id INT(20) PRIMARY KEY AUTO_INCREMENT  NOT NULL,
	username VARCHAR(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
	nickname VARCHAR(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	password VARCHAR(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
	salt VARCHAR(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
	profile_pic VARCHAR(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	create_time DATETIME ,
	update_time DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP ,
	UNIQUE KEY indx_username (username)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; `

	// 执行 SQL 语句
	_, err := db.GetMysqlClient().Exec(sql)
	if err != nil {
		fmt.Printf("create table failed: %v\n", err)
		return
	}
	fmt.Println("create table success")
}
