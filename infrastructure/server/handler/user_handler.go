package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarf/sample_clean/infrastructure/database"
	"github.com/hmarf/sample_clean/infrastructure/server/response"
	"github.com/hmarf/sample_clean/interface/controllers"
	"github.com/hmarf/sample_clean/interface/datastore"
	"github.com/hmarf/sample_clean/usecase/service"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	GetUser(http.ResponseWriter, *http.Request)
}

func NewUserHandler(db *database.MysqlData) UserHandler {
	return &userHandler{
		userController: controllers.NewUserController(
			service.NewUserService(
				datastore.NewUserRepository(db),
			),
		),
	}
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var userGetRequest controllers.UserGetRequest
	json.Unmarshal(body, &userGetRequest)
	userGetResponse, err := uh.userController.GetUser(userGetRequest.UserID)
	if err != nil {
		log.Println(err)
		return
	}
	response.Success(w, userGetResponse)
}
