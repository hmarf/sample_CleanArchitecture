package controllers

import "github.com/hmarf/sample_clean/usecase/service"

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUser(string) (*UserGetResponse, error)
	// UpdateUser(string, updateRequest *UserUpdateRequest) error
}

func NewUserController(us service.UserService) UserController {
	return &userController{userService: us}
}

type UserGetRequest struct {
	UserID string `json:"userId"`
}

type UserGetResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func (u *userController) GetUser(userID string) (*UserGetResponse, error) {
	user, err := u.userService.GetUserService(userID)

	if err != nil {
		return nil, err
	}
	return &UserGetResponse{UserID: user.UserID, Name: user.Name, CreatedAt: user.CreatedAt}, err
}
