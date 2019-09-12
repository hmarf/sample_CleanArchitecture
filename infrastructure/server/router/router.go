package router

import (
	"github.com/hmarf/sample_clean/infrastructure/server"
	"github.com/hmarf/sample_clean/infrastructure/server/handler"
)

func BootRouter(s server.Server, handler handler.RootHandler) {
	s.Post("/get/user", handler.GetUser())
	s.Post("/create/user", handler.InsertUser())
	s.Post("/delete/user", handler.DeleteUser())
}
