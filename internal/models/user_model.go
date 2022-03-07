package models

type User struct {
	ID        string  `json:"id"`
	Login     string  `json:"login"`
	Hash      *string `json:"hash"`
	Role      int     `json:"role"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
