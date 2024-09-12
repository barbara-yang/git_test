package router

import (
	"entry_task/userweb/handler"
	"net/http"
)

// InitHTTPRouter init http router
func InitHTTPRouter() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/regist", handler.Regist)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.User)
	http.HandleFunc("/edit/nickname", handler.EditNickname)
	http.HandleFunc("/edit/profilepic", handler.EditProfilePic)
	http.HandleFunc("/logout", handler.Logout)
	http.Handle("/web/img/", http.FileServer(http.Dir(".")))
	http.Handle("/web/default/", http.FileServer(http.Dir(".")))
}
