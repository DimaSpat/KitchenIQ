package routes

import (
	"fmt"
	"net/http"
)

func ApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/", baseHandler)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello over backend")
}
