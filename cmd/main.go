package main

import (
	"fmt"
	"log"

	"github.com/hmarf/sample_clean/config"
	"github.com/hmarf/sample_clean/infrastructure/database"
	"github.com/hmarf/sample_clean/infrastructure/server"
	"github.com/hmarf/sample_clean/infrastructure/server/handler"
	"github.com/hmarf/sample_clean/infrastructure/server/router"
)

func main() {
	// server, database情報を取得
	addr, databaseData := config.LoadConfig()
	// DB接続
	connectedDB := database.SetUpDB(databaseData)
	// intaractor を作成
	intaractor := handler.NewInteractor(connectedDB)
	// AppHandlerの取得
	rootHandler := intaractor.NewRootHandler()
	// Routerの起動
	serv := server.New()
	router.BootRouter(serv, rootHandler)
	// DBのClose
	defer func() {
		if err := connectedDB.DB.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	// Server Start
	fmt.Println("server start")
	serv.Start(addr)
}
