// Package webclient is a fantastic library
package webclient

import (
	"fmt"
	"io"
	"net/http"
)

func GetPage(path string, group *string) string {
	if group != nil {
		path += *group
	}

	fmt.Println("Fetching URL:", path)

	response, err := http.Get(path)
	if err != nil {
		println("Error fetching groups:", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		println("Error reading response body:", err)
	}

	return string(body)
}

// GetGroups Get all groups
func GetGroups() string {
	body := GetPage("https://attack.mitre.org/groups/", nil)
	return body
}

// GetGroup Get detailed group information
func GetGroup(group string) string {
	body := GetPage("https://attack.mitre.org/groups/", &group)
	return body
}

func GetGroupJson(group string) string {
	group = group + "/" + group + "-enterprise-layer.json"
	body := GetPage("https://attack.mitre.org/groups/", &group)
	return body
}
