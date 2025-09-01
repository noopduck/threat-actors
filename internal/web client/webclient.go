// Package webclient is a fantastic library
package webclient

import (
	"io"
	"net/http"
)

func GetGroups() string {
	response, err := http.Get("https://attack.mitre.org/groups/")
	if err != nil {
		println("Error fetching groups:", err)
	}
	defer response.Body.Close()

	// Process the response body as needed
	// For now, just return an empty string as a placeholder

	body, err := io.ReadAll(response.Body)
	if err != nil {
		println("Error reading response body:", err)
		return ""
	}

	return string(body)
}
