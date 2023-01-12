package model

type User struct {
	ID         int64   `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Password   string  `json:"password" db:"password"`
	Mail       string  `json:"mail" db:"mail"`
	Created_at *string `json:"created_at" db:"created_at"`
	Updated_at *string `json:"updated_at" db:"updated_at"`
	Deleted_at *string `json:"deleted_at" db:"deleted_at"`
}
