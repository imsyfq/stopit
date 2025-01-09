package models

type Action struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id" gorm:"not null"`
	Name          string `json:"name" gorm:"not null"`
	CurrentStreak uint   `json:"current_streak" gorm:"default:0;not null"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	User          User   `gorm:"foreignKey:UserId"`
}
