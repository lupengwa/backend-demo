package demo

import "time"

type UserEntity struct {
	Id        string    `gorm:"column:id"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (UserEntity) TableName() string {
	return "bookstore.users"
}
