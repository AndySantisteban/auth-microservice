package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
type user struct {
	email        string
	username     string
	passwordhash string
	fullname     string
	createDate   string
	role         int
}


var userList = []user{
	{
		email:        "andyjosue160720@gmail.com",
		username:     "andy",
		passwordhash: "password",
		fullname:     "Andy Santisteban",
		createDate:   "2020-07-20",
		role:         1,
	},
	{
		email:        "nico@gmail.com",
		username:     "nico",
		passwordhash: "password2",
		fullname:     "Nicolette Pacheco",
		createDate:   "2020-07-20",
		role:         0,
	},
}

func GetUserObject(email string) (user, bool) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/usuarios")

    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    results, err := db.Query("SELECT username, password FROM users")
    if err != nil {
        panic(err.Error()) 
    }

    for results.Next() {
        var tag user
        err = results.Scan(&tag.email, &tag.passwordhash)
        if err != nil {
            panic(err.Error()) 
        }
        
		if tag.email == email {
			return tag, true
	}
        log.Printf(tag.email, tag.passwordhash)
    }

	return user{}, false
}

func (u *user) ValidatePasswordHash(pswdhash string) bool {
	return u.passwordhash == pswdhash
}

func AddUserObject(email string, username string, passwordhash string, fullname string, role int) bool {
	newUser := user{
		email:        email,
		passwordhash: passwordhash,
		username:     username,
		fullname:     fullname,
		role:         role,
	}
	for _, ele := range userList {
		if ele.email == email || ele.username == username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}
