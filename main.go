package main

import (
	"fmt"
	"log"
  "os"
	"net/http"
	"./router"
)

func main() {
	r := router.Router()
	log.Println(fmt.Sprintf("Starting server on port %s...", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
