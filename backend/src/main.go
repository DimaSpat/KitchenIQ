package main

import (
	"backend/src/routes"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Note: Could not find %s file. Using system environment variables.\n", filename)
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
			log.Printf("Invalid line format, skipping: %s\n", line)
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

func main() {
	if err := loadEnv(".env"); err != nil {
		log.Fatalf("Fatal error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")

	router := http.NewServeMux()
	routes.ApiRoutes(router)

	start(router, PORT)
}

func start(router *http.ServeMux, port string) {
	log.Printf("🚀 Server running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
