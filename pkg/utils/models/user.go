package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"password"`
}

type UserToken struct {
	User  UserDetailsResponse
	Token string
}

type UserLogin struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}
