package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

const icon = "ᕙ(˘ ₃˘)ᕗ"
const title = "manly"

func titleView() string {
	title := fmt.Sprintf(
		"%s%s%s%s%s",
		"\n",
		"\n",
		icon+" "+title+" "+icon,
		"\n",
		"\n",
	)
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#F25D94")).
		PaddingLeft(4).
		Width(32)
	return titleStyle.Render(title)
}

func inputView(m model) string {
	// Configure lipgloss styles for focused and unfocused states
	focusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(30).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	unfocusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Width(30).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	// Determine which style to use based on focus
	var inputStyle lipgloss.Style
	if m.focused == INPUT {
		inputStyle = focusedStyle
	} else {
		inputStyle = unfocusedStyle
	}

	// Render the input view
	return inputStyle.Render(m.input.View())
}

func render(m model) string {
	m.list.SetHeight(20)
	return fmt.Sprintf(
		"%s%s%s%s%s",
		titleView(),
		"\n",
		inputView(m),
		"\n\n",
		m.list.View(),
	)
}
