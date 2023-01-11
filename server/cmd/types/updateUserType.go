package types

type UpdateUserRequest struct {
	ID       int64  `json:"id"`
}

type UpdateUserBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type UpdateUserResponse struct {
	
}