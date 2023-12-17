package main

import (
	router "ASimpleProgram/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is Starting.....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started at 4000 port...")
}
