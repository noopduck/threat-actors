package main

import (
	"fmt"
	"net/http"
	"threatactors/internal/parser"
	"threatactors/internal/webclient"
)

func mitreThreatGroupsHandler(w http.ResponseWriter, request *http.Request) {

	rawDocument := webclient.GetGroups()

	parsedDocument := parser.ParseHTMLTable(rawDocument)

	w.Write([]byte(parsedDocument))
}

func startAPI() {

	http.HandleFunc("/mitreThreatGroup", mitreThreatGroupsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Fatal when trying to start http listener", err.Error())
	}

}

func main() {
	// This is a placeholder main function.
	// The actual implementation would go here.
	//rawDocument := webclient.GetGroups()

	// fmt.Println(rawDocument)

	//parsedDocument := parser.ParseHTMLTable(rawDocument)

	//fmt.Println(parsedDocument)
	//
	startAPI()
}
