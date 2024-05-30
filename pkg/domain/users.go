package domain

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"phone" gorm:"unique"`
	Blocked  bool   `json:"bocked" gorm:"default:false"`
}
type Adress struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserId    uint   `json:"user_id"`
	User      User   `json:"-" foreignkey:"UserId"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Phone     string `json:"phone" gorm:"phone"`
	Pin       string `json:"pin" validate:"required"`
	Default   bool   `json:"default" gorm:"default:false"`
}
