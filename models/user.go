package models

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name" gorm:"not null"`
	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"-" gorm:"not null"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
