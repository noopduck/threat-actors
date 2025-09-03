// Package parser is a fantastic html table parser
package parser

import (
	"encoding/json"
	"strings"

	"golang.org/x/net/html"
)

type Row struct {
	ID               string `json:"ID"`
	Name             string `json:"Name"`
	AssociatedGroups string `json:"AssociatedGroups"`
	Description      string `json:"Description"`
	IDURL            string `json:"IDURL"`
	NameURL          string `json:"NameURL"`
}

type DetailsRow struct {
	Domain  string `json:"Domain"`
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Use     string `json:"Use"`
	IDURL   string `json:"IDURL"`
	NameURL string `json:"NameURL"`
}

// extractRow extracts data from a tr
func ExtractRow(tr *html.Node) *Row {
	tds := []*html.Node{}
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "td" {
			tds = append(tds, c)
		}
	}
	if len(tds) < 4 {
		return nil
	}

	// ID
	id := findFirstElement(tds[0], "a")
	idText := getText(id)
	idHref := getAttr(id, "href")

	// Name
	name := findFirstElement(tds[1], "a")
	nameText := getText(name)
	nameHref := getAttr(name, "href")

	// Associated groups
	associatedGroups := strings.Join(strings.Fields(getText(tds[2])), " ")

	// Description (plain text)
	desc := strings.Join(strings.Fields(getText(tds[3])), " ")

	return &Row{
		ID:               strings.TrimSpace(idText),
		Name:             strings.TrimSpace(nameText),
		AssociatedGroups: associatedGroups,
		Description:      desc,
		IDURL:            idHref,
		NameURL:          nameHref,
	}
}

func ExtractDetailRow(tr *html.Node) *DetailsRow {
	tds := []*html.Node{}
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "td" {
			tds = append(tds, c)
		}
	}
	if len(tds) < 4 {
		return nil
	}

	// ID
	domain := findFirstElement(tds[0], "a")
	domainText := getText(domain)
	domainHref := getAttr(domain, "href")

	// Name
	id := findFirstElement(tds[1], "a")
	idText := getText(id)
	idHref := getAttr(id, "href")

	// Associated groups
	name := strings.Join(strings.Fields(getText(tds[2])), " ")

	// Description (plain text)
	use := strings.Join(strings.Fields(getText(tds[3])), " ")

	return &DetailsRow{
		Domain:  strings.TrimSpace(domainText),
		ID:      strings.TrimSpace(idText),
		Name:    name,
		Use:     use,
		IDURL:   idHref,
		NameURL: domainHref,
	}
}

// helper functions
func getText(n *html.Node) string {
	if n == nil {
		return ""
	}
	if n.Type == html.TextNode {
		return n.Data
	}
	var out string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		out += getText(c)
	}
	return out
}

func findFirstElement(n *html.Node, tag string) *html.Node {
	if n == nil {
		return nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == tag {
			return c
		}
	}
	return nil
}

func getAttr(n *html.Node, key string) string {
	if n == nil {
		return ""
	}
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// ParseHTMLTable parses an HTML string and extracts data from tables (MITRE ATT&CK Groups)
func ParseHTMLTable[T any](htmli string, extractor func(*html.Node) *T) []byte {
	doc, err := html.Parse(strings.NewReader(htmli))
	if err != nil {
		println("Error parsing HTML:", err)
		return nil
	}

	var rows []T = make([]T, 0)

	// Find tbody's tr
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			row := extractor(n)
			if row != nil {
				rows = append(rows, *row)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	out, _ := json.MarshalIndent(rows, "", "  ")

	return out
}
