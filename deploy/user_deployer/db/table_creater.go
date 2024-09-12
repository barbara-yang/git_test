package db

import "entry_task/usersrv/dao"

// TableInit creat users table if the table not exist
func TableInit() {
	userDB := dao.UserDao{}
	userDB.CreateTable()
}
