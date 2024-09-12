package handler

import (
	"entry_task/common/module"
	"entry_task/userweb/config"
	"entry_task/userweb/service"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// User handle get user request
func User(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// get user session
		userIDCookie, err := req.Cookie("userid")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId parse fail err:%v", err)
			if err != nil {
				log.Printf("usersrv page error load fail:%v", err)
			}
			return
		}
		sessionCookie, err := req.Cookie("session")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "session expire or no auth err:%v", err)
			if err != nil {
				log.Printf("usersrv page error load fail:%v", err)
			}
			return
		}
		session := sessionCookie.Value
		id, err := strconv.Atoi(userIDCookie.Value)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userid convert err:%v", err)
			if err != nil {
				log.Printf("usersrv page error load fail:%v", err)
			}
			return
		}

		// remote call
		user, err := service.GetUser(module.User{UserID: id, Session: session})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = res.Write([]byte(fmt.Sprint("Edit username fail err: " + err.Error() + "\n")))
			if err != nil {
				log.Printf("user page load err:%s", err)
			}
			return
		}

		// page load
		t, err := template.ParseFiles(config.TEMPLATEPATH + "/user.html")
		if err != nil {
			log.Printf("usersrv page parse err: %v\n", err.Error())
		} else {
			err = t.Execute(res, user)
			if err != nil {
				log.Printf("usersrv page load err: %v\n", err.Error())
				return
			}
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}

}
