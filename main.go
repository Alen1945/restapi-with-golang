package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restTodoApp/handlers"
	"restTodoApp/postgres"

	"github.com/go-pg/pg/v9"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "belajar",
		Password: "12345678",
		Database: "db_belajar_golang",
	})
	defer DB.Close()
	app := handlers.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), app)
	if err != nil {
		log.Fatalf("cannot start server :\n%v", err)
	}
}
