package domain

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminToken struct {
	Admin   models.AdminDetailResponse
	Token   string
	Refresh string
}
