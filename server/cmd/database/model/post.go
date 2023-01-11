package model

type Post struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Descriptions string `json:"descriptions"`
	Content string `json:"content"`
	UserID int64 `json:"user_id"`
	Created_at *string `json:"created_at"`
	Updated_at *string `json:"updated_at"`
	Deleted_at *string `json:"deleted_at"`
}