package apis

import (
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// We want contract with apis.Api
type ManApi struct{}

func (m ManApi)GetList(query string) tea.Cmd {
	return func() tea.Msg {
		out, _ := exec.Command("apropos", query).Output()
		strStdout := string(out)
		// Convert string to slice for list
		results := strings.Split(strStdout, "\n")
		return ListFetchedMessage{Results: results}
	}
}

// func QuerySpecificPage(page string) string {
// 	page = stripParentheses(page)
// 	out, _ := exec.Command("man", "-P", "cat", page).Output()
// 	strStdout := string(out)
// 	// Convert string to slice for list
// 	return strStdout
// }
//
// func stripParentheses(s string) string {
// 	re := regexp.MustCompile(`\([^)]+\)`)
// 	stripped := re.ReplaceAllString(s, "")
// 	trimmed := strings.TrimSpace(stripped)
// 	return trimmed
// }
//
// func manReponseToItemList(results []string) []list.Item {
// 	// Convert list to []list.Item
// 	var list []list.Item
// 	for _, result := range results {
// 		// We also format the return man string so we can have title and desc
// 		parts := strings.Split(result, " - ")
// 		if len(parts) != 2 {
// 			continue
// 		}
// 		list = append(list, item{title: parts[0], desc: parts[1]})
// 	}
// 	return list
// }
