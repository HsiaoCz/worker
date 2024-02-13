package handlers

type UserSignup struct {
	Username   string
	Password   string
	RePassword string
	Email      string
}

type UserLogin struct {
	Username string
	Password string
	Email    string
}
