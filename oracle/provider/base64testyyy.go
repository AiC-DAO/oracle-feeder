package main_yyy

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main_test() {
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Step 1: Marshal the JSON data
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Step 2: Encode the JSON data to Base64
	base64Data := base64.StdEncoding.EncodeToString(jsonData)

	// Step 3: URL encode the Base64 string
	encodedData := url.QueryEscape(base64Data)

	// Step 4: Append the encoded Base64 string to the GET request URL
	url := fmt.Sprintf("https://example.com/api?data=%s", encodedData)

	// Send the GET request with the encoded data in the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to send GET request:", err)
		return
	}

	// Handle the response...
}
