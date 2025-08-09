package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GoldHealth Backend is running on port 8080")

	// Route /hello trả JSON cho frontend gọi
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// CORS header
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Xử lý preflight request OPTIONS
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello from GoldHealth backend!"}`)
	})

	// Route / mặc định
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to GoldHealth API")
	})

	http.ListenAndServe(":8080", nil)
}
