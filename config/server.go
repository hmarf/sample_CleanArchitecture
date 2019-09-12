package config

import "flag"

type Database struct {
	UserName string
	Password string
	Protocol string
	Host     string
	Port     string
	DBName   string
	Option   string
}

// LoadConfig はサーバー情報,DB情報を返す
func LoadConfig() (string, Database) {

	// 起動するサーバー情報
	var addr string
	flag.StringVar(&addr, "addr", ":9000", "tcp host:port to connect")
	flag.Parse()

	// 接続するDB情報
	var database Database
	database.UserName = "user"
	database.Password = "password"
	database.Protocol = "tcp"
	database.Host = "0.0.0.0"
	database.Port = "3306"
	database.DBName = "sampleDB"
	database.Option = "parseTime=true"
	return addr, database
}
