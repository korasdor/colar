package main

import (
	"github.com/gorilla/mux"

	"net/http"
	"github.com/korasdor/colar/server/handler"
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