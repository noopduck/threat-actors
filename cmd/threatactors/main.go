package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"threatactors/internal/parser"
	"threatactors/internal/webclient"
)

func apiHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	help := `
	[
  	{
  		"METHOD": "GET",
  		"endpoint": "mitreThreatGroups"
  	},
  	{
  		"METHOD": "GET",
  		"endpoint": "mitreThreatGroupDetails,
  		"query": "group=G0000"
  	}, 
  	{
  		"METHOD": "GET",
  		"endpoint": "mitreThreatGroupSearch",
  		"query": "searchTerm=word"
  	}
	]`

	w.Write([]byte(help))
}

func mitreThreatGroupsHandler(w http.ResponseWriter, request *http.Request) {
	rawDocument := webclient.GetGroups()

	parsedDocument := parser.ParseHTMLTable(rawDocument, parser.ExtractRow)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(parsedDocument))
}

func mitreThreatGroupDetailsHandler(w http.ResponseWriter, request *http.Request) {
	group := request.URL.Query().Get("group")

	rawDocument := webclient.GetGroup(group)

	parsedDocument := parser.ParseHTMLTable(rawDocument, parser.ExtractDetailRow)

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsedDocument)
}

func mitreThreatGroupSearchHandler(w http.ResponseWriter, request *http.Request) {
	searchTerm := strings.ToLower(request.URL.Query().Get("searchTerm"))

	rawDocument := webclient.GetGroups()

	parsedDocument := parser.ParseHTMLTable(rawDocument, parser.ExtractRow)

	var parsedDocumentJSON []parser.Row
	err := json.Unmarshal(parsedDocument, &parsedDocumentJSON)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	var results []parser.Row
	for _, row := range parsedDocumentJSON {
		if strings.Contains(strings.ToLower(row.ID), searchTerm) ||
			strings.Contains(strings.ToLower(row.Name), searchTerm) ||
			strings.Contains(strings.ToLower(row.Description), searchTerm) {

			results = append(results, row)
		}
	}

	response, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func startAPI() {
	http.HandleFunc("/", apiHandler)
	http.HandleFunc("/mitreThreatGroups", mitreThreatGroupsHandler)
	http.HandleFunc("/mitreThreatGroupDetails", mitreThreatGroupDetailsHandler)
	http.HandleFunc("/mitreThreatGroupSearch", mitreThreatGroupSearchHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Fatal when trying to start http listener", err.Error())
	}
}

func main() {
	// This is a placeholder main function.
	// The actual implementation would go here.
	// rawDocument := webclient.GetGroups()

	// fmt.Println(rawDocument)

	// parsedDocument := parser.ParseHTMLTable(rawDocument)

	//fmt.Println(parsedDocument)
	//
	startAPI()
}
