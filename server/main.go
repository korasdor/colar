package main

import (
	"github.com/gorilla/mux"

	"fmt"
	"net/http"
	"src/handler"
)

func main() {


	r := mux.NewRouter()
	r.HandleFunc("/", 	handler.IndexHandler)
	//r.HandleFunc("/static/media/{file}", handler.ServeStaticFiles)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", "s")
}
