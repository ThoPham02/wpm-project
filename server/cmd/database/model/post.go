package model

type Post struct {
	ID           int64   `json:"id" db:"id"`
	Title        string  `json:"title" db:"title"`
	Descriptions string  `json:"descriptions" db:"descriptions"`
	Content      string  `json:"content" db:"content"`
	UserID       int64   `json:"user_id" db:"user_id"`
	Created_at   *string `json:"created_at" db:"created_at"`
	Updated_at   *string `json:"updated_at" db:"updated_at"`
	Deleted_at   *string `json:"deleted_at" db:"deleted_at"`
}
