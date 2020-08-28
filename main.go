package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/simplifyd-examples/golang-mongodb-todo-app/router"
)

func main() {
	r := router.Router()
	log.Println(fmt.Sprintf("Starting server on port %s...", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
