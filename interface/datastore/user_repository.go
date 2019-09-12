package datastore

import (
	"log"
	"time"

	"github.com/hmarf/sample_clean/domain"
	"github.com/hmarf/sample_clean/usecase/repository"
)

type userRepository struct {
	db ConnectedSql
}

func NewUserRepository(db ConnectedSql) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Insert(userID string, name string, createdAt time.Time) (err error) {
	_, err = u.db.Exec("INSERT INTO user(user_id,name,createdAt) VALUES(?,?,?)", userID, name, createdAt)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *userRepository) Select(userID string) (user domain.User, err error) {
	row := u.db.QueryRow("select * from user where user_id=?", userID)
	err = row.Scan(&user.UserID, &user.Name, &user.CreatedAt)
	if err != nil {
	}
	return
}

func (u *userRepository) Delete(userID string) (err error) {
	_, err = u.db.Exec("DELETE FROM user WHERE user_id=?", userID)
	return
}
