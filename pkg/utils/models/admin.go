package models

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=15"`
}

type AdminDetailResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserdetailsAtAdmin struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Blocked string `json:"blocked"`
}
