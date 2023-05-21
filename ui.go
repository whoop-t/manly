package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

const icon = "ᕙ(˘ ₃˘)ᕗ"
const title = "manly"

func titleView(m model) string {
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
		Width(m.windowWidth)
	return titleStyle.Render(title)
}

func inputView(m model) string {
	// Offset specific to input
	const widthOffset = 2
	inputWidth := m.windowWidth/3 - widthOffset
	// Configure lipgloss styles for focused and unfocused states
	focusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(inputWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	unfocusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Width(inputWidth).
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

func listView(m model) string {
	// listHeight := calculateListHeight(m)
	m.list.SetHeight(18)
	m.list.SetWidth(m.windowWidth)
	return m.list.View()
}

func pageView(m model) string {
	m.page.Style.Width(m.windowWidth).Height(m.windowHeight)
	return m.page.View()
}

func render(m model) string {
	if m.focused != PAGE {
		return fmt.Sprintf(
			"%s%s%s%s%s",
			titleView(m),
			"\n",
			inputView(m),
			"\n\n",
			listView(m),
		)
	}

	return pageView(m)
}
