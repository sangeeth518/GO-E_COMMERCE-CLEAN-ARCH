package domain

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"phone" gorm:"unique"`
	Blocked  bool   `json:"bocked" gorm:"default:false"`
}
