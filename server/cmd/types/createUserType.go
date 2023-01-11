package types

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type CreateUserResponse struct {
}
