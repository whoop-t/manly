package main

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)


func newPage(content string, m model) viewport.Model {
	page := viewport.New(m.windowWidth, m.windowHeight)
	page.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.windowWidth),
	)
	if err != nil {
		log.Fatal(err)
	}

	str, err := renderer.Render(content)
	if err != nil {
		log.Fatal(err)
	}

	page.SetContent(str)

	return page
}
