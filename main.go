package main

import (
	"fmt"
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
	api := NewAPIServer(":8080", store)
	api.Serve()

	fmt.Println(len([]int{1, 2, 3, 4, 5})) // 5

	// fmt.Println(fruitSlice[1:3]) // [StartIndex:EndIndex-1] ==> [1:3] ==> [Orange Grape] ==> 3-1 = 2 elements

	// // empty slice
	// var fruitSlice2 []string

	// // append ==> add to slice
	// fruitSlice2 = append(fruitSlice2, "Apple", "Orange", "Grape", "Cherry")

	// fmt.Println(fruitSlice2)

}
