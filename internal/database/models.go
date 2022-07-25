package database

import "time"

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

type User struct {
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	UserEmail string    `json:"user_email"`
}
