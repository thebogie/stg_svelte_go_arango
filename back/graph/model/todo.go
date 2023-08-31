package model

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      string `json:"userID"`
	User        *User  `json:"user" gorm:"foreignKey:UserID"`
}
