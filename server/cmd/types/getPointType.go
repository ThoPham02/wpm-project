package types

type GetPointRequest struct {
}

type Point struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	PostID     int64  `json:"post_id"`
	Wpm        int64  `json:"wpm"`
	Created_at *string `json:"created_at"`
}

type GetPointResponse struct {
	Data []Point `json:"data"`
}
