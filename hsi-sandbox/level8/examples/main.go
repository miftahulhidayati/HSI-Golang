package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    // Example of making an API request with JWT authentication

    // Step 1: Create a user
    user := User{
        Username: "testuser",
        Password: "testpassword",
    }

    // Step 2: Convert user to JSON
    userJson, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error marshalling user:", err)
        return
    }

    // Step 3: Make a POST request to authenticate the user
    resp, err := http.Post("http://example.com/api/auth", "application/json", bytes.NewBuffer(userJson))
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Step 4: Check if authentication was successful
    if resp.StatusCode == http.StatusOK {
        fmt.Println("Authentication successful!")
        // Here you would typically read the JWT token from the response
    } else {
        fmt.Println("Authentication failed with status:", resp.Status)
    }
}