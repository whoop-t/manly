package main

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)
type item string

func (i item) FilterValue() string { return "" }

func queryManPages(query string) []list.Item {
	out, _ := exec.Command("man", "-k", query).Output()
	strStdout := string(out)
	// Convert string to slice for list
	results := strings.Split(strStdout, "\n")
	
	// Convert list to []list.Item
	var list []list.Item
	for _, result := range results {
		list = append(list, item(result))
	}
	return list
}
