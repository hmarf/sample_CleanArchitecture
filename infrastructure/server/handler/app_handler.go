package handler

import (
	"net/http"

	"github.com/hmarf/sample_clean/infrastructure/database"
)

type interactor struct {
	db *database.MysqlData
}

type Interactor interface {
	NewRootHandler() RootHandler
}

func NewInteractor(db *database.MysqlData) Interactor {
	return &interactor{db: db}
}

type rootHandler struct {
	userHandler UserHandler
}

type RootHandler interface {
	GetUser() http.HandlerFunc
}

func (i *interactor) NewRootHandler() RootHandler {
	return &rootHandler{
		userHandler: NewUserHandler(i.db),
	}
}

func (uh *rootHandler) GetUser() http.HandlerFunc {
	return uh.userHandler.GetUser
}
