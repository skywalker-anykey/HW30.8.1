package main

import (
	"fmt"
	"log"
	"pgsql/pkg/storage"
	"pgsql/pkg/storage/postgres"
)

var db storage.Interface

const (
	dbUser = "sandbox"
	dbPass = "sandbox"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "tasks"
)

func main() {
	var err error
	// В случае продуктового использования пароль "прячем"
	//pwd := os.Getenv("dbbass")
	//if pwd == "" {
	//	os.Exit(1)
	//}

	connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err = postgres.New(connect)
	if err != nil {
		log.Fatal(err)
	}

	// Тестирование
	// db = memdb.DB{}

	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)

}
