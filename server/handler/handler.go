package handler

import (
	"net/http"
	"fmt"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", "s")
}
