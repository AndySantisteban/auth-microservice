package data

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
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
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
