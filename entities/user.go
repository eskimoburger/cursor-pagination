package entities

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Cursor struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
}
