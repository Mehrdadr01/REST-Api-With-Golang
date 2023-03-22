package main

import (
	"fmt"
	"net/http"
	"os"

	"handlers"
)

func main() {

	// https://pkg.go.dev/net/http#HandleFunc
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		os.Exit(1)
	}
}

//////////////////////////  rootHandlers   ///////////////////////////////////////////
