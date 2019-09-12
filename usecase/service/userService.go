package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hmarf/sample_clean/domain"
	"github.com/hmarf/sample_clean/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
}

type UserService interface {
	GetUserService(string) (domain.User, error)
	InsertUserService(string) (domain.User, error)
	DeleteUserService(string) error
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{UserRepository: ur}
}

func (u *userService) InsertUserService(name string) (user domain.User, err error) {

	// userID　を作成する
	userID, err := uuid.NewRandom()
	if err != nil {
		return
	}

	nowTime := time.Now()
	user.UserID = userID.String()
	user.Name = name
	user.CreatedAt = nowTime.String()
	err = u.UserRepository.Insert(user.UserID, user.Name, nowTime)
	return
}

func (u *userService) GetUserService(userID string) (domain.User, error) {
	user, err := u.UserRepository.Select(userID)
	return user, err
}

func (u *userService) DeleteUserService(userID string) error {
	return errors.New("error")
}
