package model

type Point struct {
	ID         int64   `json:"id"`
	UserID     string  `json:"user_id"`
	PostID     string  `json:"post_id"`
	Wpm        string  `json:"wpm"`
	Created_at *string `json:"created_at"`
	Updated_at *string `json:"updated_at"`
	Deleted_at *string `json:"deleted_at"`
}
