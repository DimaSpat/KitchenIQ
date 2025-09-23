package main

import (
	"backend/src/routes"
	"fmt"
	"net/http"
)

const PORT = ":3000"

func main() {
	router := http.NewServeMux()
	routes.ApiRoutes(router)

	start(router)
}

func start(router *http.ServeMux) {
	fmt.Printf("🚀 Server running on http://localhost%s\n", PORT)
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
