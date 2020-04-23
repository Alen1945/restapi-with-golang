package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restTodoApp/handlers"
)

func main() {
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
