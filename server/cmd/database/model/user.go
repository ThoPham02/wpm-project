package model

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Mail string `json:"mail"`
	Created_at *string `json:"created_at"`
	Updated_at *string `json:"updated_at"`
	Deleted_at *string `json:"deleted_at"`
}