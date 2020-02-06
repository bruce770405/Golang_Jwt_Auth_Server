package login

type UserCredentials struct {
	UserName string
	PxssCode string
}

type Token struct {
	Token string
}

type User struct {
	UserName string
	PxssCode string
	Email    string
}
