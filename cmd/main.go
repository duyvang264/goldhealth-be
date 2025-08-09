package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(("GoldHealth Backend is running on port 8080"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to GoldHealth API")
	})
	http.ListenAndServe(":8080", nil)
}
