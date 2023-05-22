package pager

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/whoop-t/manly/internal/apis"
)

type Model struct {
	Page      viewport.Model
	IsPageSet bool
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.IsPageSet {
			if msg.String() == "j" {
				m.Page.LineDown(2)
			}
			if msg.String() == "k" {
				m.Page.LineUp(2)
			}
			if msg.String() == "q" || msg.String() == "esc" {
				m.IsPageSet = false
			}
		}
	case apis.ShowPageMessage:
		fmt.Println("pager stuff")
		m.Page = NewPage(msg.Result)
		m.IsPageSet = true
		m.Page, cmd = m.Page.Update(msg)
	}
	return m, cmd
}

func NewPage(content string) viewport.Model {
	const width = 90

	page := viewport.New(width, 25)
	page.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
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
