package types

type UpdatePostRequest struct {
	ID int64 `uri:"id"`
}

type UpdatePostBody struct {
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Content      string `json:"content"`
}

type UpdatePostResponse struct {
}
