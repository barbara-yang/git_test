package handler

import (
	"entry_task/common/module"
	"entry_task/userweb/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// EditProfilePic profilePic edit handler
func EditProfilePic(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// form parse
		req.ParseMultipartForm(32 << 20)

		user := module.User{}

		// get user session userId
		userIDCookie, err := req.Cookie("userid")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId parse err:%v", err)
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
				log.Printf("user page error load fail:%v", err)
			}
			return
		}
		user.Session = sessionCookie.Value
		user.UserID, err = strconv.Atoi(userIDCookie.Value)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId convert err:%v", err)
			if err != nil {
				log.Printf("user page error load fail:%v", err)
			}
			return
		}

		// file receive handle
		file, handler, err := req.FormFile("pic")
		if err != nil {
			log.Println(err)
			user.ProfilePic = ""
		} else {
			user.ProfilePic = service.FileUpload(file, handler)
		}

		if user.ProfilePic == "" {
			http.Redirect(res, req, "/user", http.StatusFound)
			return
		}

		// edit profilePic remote call
		_, err = service.EditUserProfilePic(user)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "remote call err: %s", err)
			if err != nil {
				log.Printf("user error page load err:%v\n", err)
			}
			return
		}
		http.Redirect(res, req, "/user", http.StatusFound)
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}
}

// EditNickname handle edit nickname request
func EditNickname(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// parse form
		nickname := req.FormValue("nickname")

		// get user session
		sessionCookie, err := req.Cookie("session")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "session expire or no auth err:%v", err)
			if err != nil {
				fmt.Errorf("user page error load fail:%v", err)
			}
			return
		}
		useridCookie, err := req.Cookie("userid")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId parse err:%v", err)
			if err != nil {
				log.Printf("user page error load fail:%v", err)
			}
			return
		}
		session := sessionCookie.Value
		id, err := strconv.Atoi(useridCookie.Value)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprintf(res, "userId convert err:%v", err)
			if err != nil {
				log.Fatalf("user page error load fail:%v", err)
			}
			return
		}

		// remote call
		_, err = service.EditNickname(module.User{Nickname: nickname, UserID: id, Session: session})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = res.Write([]byte(fmt.Sprint("Edit username fail err: " + err.Error() + "\n")))
			if err != nil {
				log.Printf("user page load err:%s", err)
			}
			return
		}

		http.Redirect(res, req, "/user", http.StatusFound)

		// page load
	} else {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "illegal access")
	}

}
