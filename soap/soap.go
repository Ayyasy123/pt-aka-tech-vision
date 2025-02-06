package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type UserDetails struct {
	XMLName xml.Name `xml:"UserDetails"`
	Name    string   `xml:"Name"`
	Email   string   `xml:"Email"`
}

func main() {
	resp, err := http.Get("http://example.com/soap-service")
	if err != nil {
		fmt.Println("Error fetching SOAP response:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var user UserDetails
	err = xml.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}
