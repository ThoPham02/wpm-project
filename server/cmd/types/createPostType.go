package types

type CreatePostRequest struct {
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Content      string `json:"content"`
	UserID       int64  `json:"user_id"`
}

type CreatePostResponse struct {
}
