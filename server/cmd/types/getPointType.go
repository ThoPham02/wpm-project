package types

type GetPointRequest struct {
}

type Point struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	PostID     string  `json:"post_id"`
	Wpm        string  `json:"wpm"`
	Created_at *string `json:"created_at"`
}

type GetPointResponse struct {
	Data []Point `json:"data"`
}
