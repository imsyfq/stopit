package models

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name" gorm:"not null"`
	ApiToken  *string `gorm:"index:api_token_index"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
