package delivery

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type changePasswordInput struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

type signUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signResponse struct {
	Message string `json:"message"`
}

type signInResponse struct {
	Token string `json:"token"`
}
