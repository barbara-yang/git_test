package handler

import (
	"entry_task/userweb/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// register handle register request
func Register(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")
		user, err := service.Register(username, password)
		if err != nil {
			fmt.Printf("register err: %s user:%s password:%s\n", err, username, password)
			res.WriteHeader(http.StatusUnauthorized)
			_, err = res.Write([]byte(fmt.Sprint("register fail err: " + err.Error() + "\n")))
			if err != nil {
				log.Printf("register page load err:%s", err)
			}
			return
		}
		// cookie set sessionId
		http.SetCookie(res, &http.Cookie{
			Name:     "session",
			Value:    user.Session,
			Expires:  time.Now().Add(3600 * time.Second),
			HttpOnly: true,
		})

		http.SetCookie(res, &http.Cookie{
			Name:     "userid",
			Value:    strconv.Itoa(user.UserID),
			Expires:  time.Now().Add(3600 * time.Second),
			HttpOnly: true,
		})
		// page load
		http.Redirect(res, req, "/user", http.StatusFound) //register成功后转到login界面
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}

}
