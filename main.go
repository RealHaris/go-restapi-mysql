package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPass,
		Net:                  "tcp",
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewAPIServer(":3000", store)

	// print ip address

	api.Serve()

}
