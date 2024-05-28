package handlers

import (
	"fmt"
	"net/http"
)

func HelloWorld(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(rw, req)
		return
	}
	fmt.Fprintf(rw, "Hello, World!")
}
