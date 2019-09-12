package controllers

import (
	"fmt"

	"github.com/hmarf/sample_clean/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUser(string) (*UserGetResponse, error)
	InsertUser(string) (*UserInsertResponse, error)
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

type UserInsertRequest struct {
	Name string `json:"name"`
}

type UserInsertResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func (u *userController) InsertUser(name string) (*UserInsertResponse, error) {
	fmt.Println("insertService")
	user, err := u.userService.InsertUserService(name)
	if err != nil {
		return nil, err
	}
	return &UserInsertResponse{UserID: user.UserID, Name: user.Name, CreatedAt: user.CreatedAt}, err
}

type UserDeleteRequest struct {
	UserID string `json:"userId"`
}

type UserDeleteResponse struct {
	Result string `json:"results"`
}

func (u *userController) DeleteUser(name string) (*UserDeleteResponse, error) {
	err := u.userService.DeleteUserService(name)
	if err != nil {
		return &UserDeleteResponse{Result: "failure"}, err
	}
	return &UserDeleteResponse{Result: "success"}, err
}
