package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// Exercise 1: Create a simple API that returns a greeting message.
func greetingHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "Hello, welcome to the Golang API!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Exercise 2: Implement JWT authentication for a protected route.
func protectedHandler(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")
    if token == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    // Here you would normally validate the token
    response := map[string]string{"message": "You have accessed a protected route!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/greet", greetingHandler)
    http.HandleFunc("/protected", protectedHandler)
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}