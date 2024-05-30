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
	Phone string `json:"phone"`
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

type UserSigninResponse struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Changepassword struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_newpassword"`
}

type AddAdress struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Phone     string `json:"phone" validate:"require"`
	Pin       string `json:"pin" validate:"required"`
}
