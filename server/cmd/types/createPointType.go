package types

type CreatePointRequest struct {
	UserID     int64  `json:"user_id"`
	PostID     int64  `json:"post_id"`
	Wpm        int64  `json:"wpm"`
}

type CreatePointResponse struct {
	
}