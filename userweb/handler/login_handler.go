package handler

import (
	"entry_task/userweb/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Login handle Login request and load page for user

func Login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")
		user, err := service.Login(username, password)
		if err != nil {
			log.Printf("login err: %s user:%s password:%s\n", err, username, password)
			res.WriteHeader(http.StatusUnauthorized)
			_, writeErr := res.Write([]byte(fmt.Sprintf("login failed: %s\n", err.Error())))
			if writeErr != nil {
				log.Printf("error writing response: %s", writeErr)
			}
			return
		}

		// Set session and user ID cookies
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

		// Redirect to user page
		http.Redirect(res, req, "/user", http.StatusFound)
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}
}
