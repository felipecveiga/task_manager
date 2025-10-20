package model

import "time"

type User struct {
	ID        int        `json:"id" gorm:"AUTO_INCREMENT:primaryKey"`
	Name      string     `json:"nome"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

type Status string

const (
	Done       Status = "done"
	Pending    Status = "pending"
	InProgress Status = "in-progress"
)

type Task struct {
	ID          int        `json:"id" gorm:"AUTO_INCREMENT:primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      Status     `json:"status" gorm:"type:enum('done','pending','in-progress');default:'pending'"`
	UserID      int        `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}
