package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// https://pkg.go.dev/net/http#HandleFunc
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		os.Exit(1)
	}
}

//////////////////////////  rootHandlers   ///////////////////////////////////////////
func rootHandler(_writer http.ResponseWriter, _req *http.Request) {
	if _req.URL.Path != "/" {
		_writer.WriteHeader(http.StatusNotFound)
		_writer.Write([]byte("item not Found\n"))
		return
	}

	_writer.WriteHeader(http.StatusOK)
	_writer.Write([]byte("running the API v1.0\n"))
	// here we didn't check to see if the route is root("/") or not
	// so we get same response for localhost:11111/hello or localhost:11111/anything
	// so we first check to see if the rote is / or not
}
