package handler

import (
	"entry_task/common/module"
	"entry_task/userweb/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Logout handle logout request for user
func Logout(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		// get user session
		userIDCookie, err := req.Cookie("userid")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId parse  err:%v", err)
			if err != nil {
				log.Printf("user page error load fail:%v", err)
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
			_, err = fmt.Fprintf(res, "userId convert err:%v", err)
			if err != nil {
				log.Printf("usersrv page error load fail:%v", err)
			}
			return
		}

		// remote call
		err = service.Logout(module.User{UserID: id, Session: session})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = res.Write([]byte(fmt.Sprint("logout fail err: " + err.Error() + "\n")))
			if err != nil {
				log.Printf("logout err :%s", err)
			}
			return
		}

		// cookie 清除

		http.SetCookie(res, &http.Cookie{
			Name:    "session",
			MaxAge:  -1,
			Expires: time.Now().Add(-100 * time.Hour),
		})
		http.SetCookie(res, &http.Cookie{
			Name:    "userid",
			MaxAge:  -1,
			Expires: time.Now().Add(-100 * time.Hour),
		})

		http.Redirect(res, req, "/", http.StatusFound) //跳转到index页面
		// page load
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}

}
