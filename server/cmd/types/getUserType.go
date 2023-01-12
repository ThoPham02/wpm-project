package types

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	CreatedAt *string `json:"created_at"`
}

type GetUserRequest struct {
	ID int64 `json:"id"`
}

type GetUserResponse struct {
	Data []User `json:"data"`
}
