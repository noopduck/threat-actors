package main

import (
	"fmt"
	"threatactors/internal/parser"
	"threatactors/internal/webclient"
)

func main() {
	// This is a placeholder main function.
	// The actual implementation would go here.
	rawDocument := webclient.GetGroups()

	// fmt.Println(rawDocument)

	parsedDocument := parser.ParseHTMLTable(rawDocument)

	fmt.Println(parsedDocument)
}
