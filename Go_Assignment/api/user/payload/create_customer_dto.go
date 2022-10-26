package customer_api_payload

type CreateUserRequest struct {
	Username    string `json:"userName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Address     string `json:"address" binding:"required,max=200"`
	Email       string `json:"email" binding:"required,max=50"`
	PhoneNumber string `json:"phoneNumber" binding:"required,min=0,max=12"`
}

type CreateUserResponse struct {
	Status    string `json:"status"`
	Username  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
}
