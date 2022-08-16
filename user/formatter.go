package user

type UserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ResponseFormatterUser(user User) UserResponse {
	response := UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return response
}

func ResponseFormatterUsers(user []User) []UserResponse {
	var response []UserResponse
	for _, value := range user {
		response = append(response, ResponseFormatterUser(value))
	}
	return response
}
