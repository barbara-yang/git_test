package handler

import (
	"entry_task/common/module"
	"entry_task/userweb/config"
	"entry_task/userweb/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Index load index page for user
func Index(res http.ResponseWriter, req *http.Request) {
	// get user session
	if req.Method == "GET" {
		userIDCookie, err := req.Cookie("userid") //userid的cokkie
		if err != nil {
			loadIndex(res)
			return
		}
		sessionCookie, err := req.Cookie("session") //session的cokkie
		if err != nil {
			loadIndex(res)
			return
		}
		session := sessionCookie.Value
		id, err := strconv.Atoi(userIDCookie.Value)
		auth, err := service.Index(module.User{UserID: id, Session: session})
		if err != nil {
			loadIndex(res)
			return
		}
		if auth {
			http.Redirect(res, req, "/user", http.StatusFound) //认证成功之后直接跳转到user
			return
		}
		loadIndex(res)

	}

}

func loadIndex(res http.ResponseWriter) {
	t, err := template.ParseFiles(config.TEMPLATEPATH + "/index.html")
	if err != nil {
		log.Println("html parse err: " + err.Error())
	} else {
		err := t.Execute(res, nil)
		if err != nil {
			log.Printf("index page load err:%v\n", err)
		}
	}
}

// Index load index page for user
func Regist(res http.ResponseWriter, req *http.Request) {
	// get user session
	if req.Method == "GET" {
		userIDCookie, err := req.Cookie("userid")
		if err != nil {
			loadRegist(res)
			return
		}
		sessionCookie, err := req.Cookie("session")
		if err != nil {
			loadRegist(res)
			return
		}
		session := sessionCookie.Value
		id, err := strconv.Atoi(userIDCookie.Value)
		auth, err := service.Index(module.User{UserID: id, Session: session})
		if err != nil {
			loadRegist(res)
			return
		}
		if auth {
			http.Redirect(res, req, "/user", http.StatusFound)
			return
		}
		loadRegist(res)

	}

}

func loadRegist(res http.ResponseWriter) {
	t, err := template.ParseFiles(config.TEMPLATEPATH + "/register.html") //加载页面
	if err != nil {
		log.Println("html parse err: " + err.Error())
	} else {
		err := t.Execute(res, nil)
		if err != nil {
			log.Printf("register page load err:%v\n", err)
		}
	}
}
