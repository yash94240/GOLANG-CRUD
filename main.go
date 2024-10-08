package main

import (
	"fmt"
	"mongosetup/routers"
	"net/http"
)

func main() {
	fmt.Println("Greetings!")
	r := routers.Router()
	http.ListenAndServe(":8080", r)
}
