package model

type Point struct {
	ID         int64   `json:"id" db:"id"`
	UserID     int64  `json:"user_id" db:"user_id"`
	PostID     int64  `json:"post_id" db:"post_id"`
	Wpm        int64  `json:"wpm" db:"wpm"`
	Created_at *string `json:"created_at" db:"created_at"`
	Updated_at *string `json:"updated_at" db:"updated_at"`
	Deleted_at *string `json:"deleted_at" db:"deleted_at"`
}
