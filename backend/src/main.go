package main

import (
	"backend/src/routes"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := loadEnv(".env"); err != nil {
			fmt.Printf("Fatal error loading .env file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("\033[33mNote: Running in production mode. Not loading .env file.\033[0m")
	}

	PORT := ":" + os.Getenv("PORT")

	router := http.NewServeMux()
	routes.ApiRoutes(router)

	handler := CORS(router)

	start(handler, PORT)
}

func start(handler http.Handler, port string) {
	fmt.Printf("🚀 Server running on http://localhost%s%s\n", port, "/")
	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}

func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Note: Could not find %s file. Using system environment variables.\n", filename)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid line format, skipping: %s\n", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading %s: %w", filename, err)
	}

	return nil
}

func CORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	}
}
