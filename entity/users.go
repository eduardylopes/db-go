package entity

import "time"

type User struct {
	ID        int        `gorm:"primary key" json:"id"`
	Nickname  string     `db:"nickname" json:"nickname"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
