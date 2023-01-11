package types

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Mail string `json:"mail"`
}

type GetUserRequest struct {
	ID int64 `json:"id"`
}

type GetUserResponse struct {
	UserResponse []User `json:"data_response"`
}