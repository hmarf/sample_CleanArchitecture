package repository

import (
	"time"

	"github.com/hmarf/sample_clean/domain"
)

//依存関係の逆転の法則
type UserRepository interface {
	Insert(string, string, time.Time) error
	Select(string) (domain.User, error)
	Delete(string) error
}
