package types

type CreatePointRequest struct {
	UserID     string  `json:"user_id"`
	PostID     string  `json:"post_id"`
	Wpm        string  `json:"wpm"`
}

type CreatePointResponse struct {
	
}