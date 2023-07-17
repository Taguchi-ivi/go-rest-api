package model

import "time"


type Tweet struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId uint `json:"user_id" gorm:"not null"`

}

type TweetResponse struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}