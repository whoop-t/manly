package apis

import (
	"os/exec"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// We want contract with apis.Api
type ManApi struct{}

func (m ManApi) GetList(query string) tea.Cmd {
	return func() tea.Msg {
		out, _ := exec.Command("apropos", query).Output()
		strStdout := string(out)
		// Convert string to slice for list
		results := strings.Split(strStdout, "\n")
		return ListFetchedMessage{Results: results}
	}
}

func (m ManApi) ShowPage(page string) tea.Cmd {
	return func() tea.Msg {
		page = stripParentheses(page)
		// Cat return the text, not the default pager
		// We want to pipe to our own pager
		out, _ := exec.Command("man", "-P", "cat", page).Output()
		strStdout := string(out)
		return ShowPageMessage{Result: strStdout}
	}
}

func stripParentheses(s string) string {
	re := regexp.MustCompile(`\([^)]+\)`)
	stripped := re.ReplaceAllString(s, "")
	trimmed := strings.TrimSpace(stripped)
	return trimmed
}
