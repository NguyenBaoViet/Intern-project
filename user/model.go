package user

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Address string    `json:"address"`
}
